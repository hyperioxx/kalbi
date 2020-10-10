package transaction

import (
	"github.com/KalbiProject/Kalbi/log"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/transport"
	"github.com/sirupsen/logrus"
	"sync"
)

//NewTransactionManager returns a new TransactionManager
func NewTransactionManager() *TransactionManager {
	txmng := new(TransactionManager)
	txmng.TX = make(map[string]Transaction)
	txmng.txLock = &sync.RWMutex{}

	return txmng
}

//TransactionManager handles SIP transactions
type TransactionManager struct {
	TX              map[string]Transaction
	RequestChannel  chan Transaction
	ResponseChannel chan Transaction
	ListeningPoint  transport.ListeningPoint
	txLock          *sync.RWMutex
}

// Handle runs TransManager
func (tm *TransactionManager) Handle(message *message.SipMsg) {

	if message.Req.StatusCode != nil {
		log.Log.Info("Client transaction")

		tx, exists := tm.FindTransaction(message)
		if exists {
			log.Log.Info("Client Transaction already exists")
		} else {
			tx = tm.NewClientTransaction(message)
		}

		tx.Receive(message)
		tm.ResponseChannel <- tx

	} else if message.Req.Method != nil {
		log.Log.Info("Server transaction")
		tx, exists := tm.FindTransaction(message)

		if exists {
			log.Log.Info("Server Transaction already exists")

		} else {
			tx = tm.NewServerTransaction(message)
			tm.PutTransaction(tx)
		}

		tx.Receive(message)
		tm.RequestChannel <- tx
	}

}



func (tm *TransactionManager) FindTransaction(msg *message.SipMsg) (Transaction, bool) {
	//key := tm.MakeKey(*msg)
	tm.txLock.RLock()
	tx, exists := tm.TX[string(msg.Via[0].Branch)]
	tm.txLock.RUnlock()
	return tx, exists
}

func (tm *TransactionManager) PutTransaction(tx Transaction) {
	tm.txLock.Lock()
	tm.TX[tx.GetBranchId()] = tx
	tm.txLock.Unlock()
}

func (tm *TransactionManager) DeleteTransaction(branch string) {
	log.Log.Info("Deleting transaction with ID: " + branch)
	log.Log.WithFields(logrus.Fields{"transactions": len(tm.TX)}).Debug("Current transaction count before DeleteTransaction() is called")
	tm.txLock.Lock()
	delete(tm.TX, branch)
	tm.txLock.Unlock()
	log.Log.WithFields(logrus.Fields{"transactions": len(tm.TX)}).Debug("Current transaction count after DeleteTransaction() is called")
}


func (tm *TransactionManager) MakeKey(msg message.SipMsg) string {

	key := string(msg.Via[0].Branch)
	var _method string
	if msg.Req.Method != nil {
         if string(msg.Req.Method) == method.ACK {
               _method = method.INVITE
		 } else {
			 _method = string(msg.Req.Method)
		 }
	} else {
		_method = string(msg.Cseq.Method)
	}

	key += _method
    return key 
}

func (tm *TransactionManager) NewClientTransaction(msg *message.SipMsg) *ClientTransaction {

	tx := new(ClientTransaction)
	tx.SetListeningPoint(tm.ListeningPoint)
	tx.TransManager = tm

	tx.InitFSM(msg)

	tx.BranchID = GenerateBranchId()

	tx.Origin = msg

	tm.PutTransaction(tx)
	return tx

}

func (tm *TransactionManager) NewServerTransaction(msg *message.SipMsg) *ServerTransaction {

	tx := new(ServerTransaction)
	tx.SetListeningPoint(tm.ListeningPoint)

	tx.TransManager = tm

	tx.InitFSM(msg)

	tx.BranchID = string(msg.Via[0].Branch)
	tx.Origin = msg

	return tx

}
