package transaction

import (
	"sync"
	"Kalbi/pkg/log"
	"Kalbi/pkg/sip/message"
	"Kalbi/pkg/transport"
	"github.com/sirupsen/logrus"	
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
	TX         map[string] Transaction
	RequestChannel    chan Transaction
	ResponseChannel   chan Transaction
	ListeningPoint    transport.ListeningPoint
	txLock            *sync.RWMutex

}



// Start runs TransManager
func (tm *TransactionManager) Handle(message *message.SipMsg)  {
	
		if  message.Req.StatusCode != nil {
			log.Log.Info("Client transaction")

		    tx, exists := tm.FindTransaction(string(message.Via[0].Branch))
			if(exists){
				log.Log.Info("Client Transaction aready exists")
			}else{
				tx = tm.NewClientTransaction(message)
				tx.SetListeningPoint(tm.ListeningPoint)  
			}


            tx.Receive(message)
			tm.ResponseChannel <- tx
			
		} else if message.Req.Method != nil {
			log.Log.Info("Server transaction")
			tx, exists := tm.FindTransaction(string(message.Via[0].Branch))

			if(exists){
				log.Log.Info("Server Transaction aready exists")
				
			}else{
				tx = tm.NewServerTransaction(message)
				tx.SetListeningPoint(tm.ListeningPoint)
				tm.PutTransaction(tx)
			}

			tx.Receive(message)
			tm.RequestChannel <- tx 
		}	
        
}


func (tm *TransactionManager) FindTransaction(branch string) (Transaction, bool) {
	tx , exists := tm.TX[branch]
	return tx , exists
}

func (tm *TransactionManager) PutTransaction(tx Transaction) {
	tm.txLock.Lock()
	tm.TX[string(tx.GetBranchId())] = tx
    tm.txLock.Unlock()
}


func (tm *TransactionManager) DeleteTransaction(branch string){
	log.Log.WithFields(logrus.Fields{"transactions": len(tm.TX)}).Info("Current transaction count before DeleteTransaction() is called")
	tm.txLock.Lock()
	delete(tm.TX, branch)
	tm.txLock.Unlock()
	log.Log.WithFields(logrus.Fields{"transactions": len(tm.TX)}).Info("Current transaction count after DeleteTransaction() is called")
}



func (tm *TransactionManager) NewClientTransaction(msg *message.SipMsg) *ClientTransaction {

	tx := new(ClientTransaction)
	tx.TransManager = tm

	tx.InitFSM(msg)


	tx.BranchID = GenerateBranchId()

	tx.Origin = msg
	
	tm.txLock.Lock()
	tm.TX[string(tx.BranchID)] = tx
    tm.txLock.Unlock()
	return tx

}

func (tm *TransactionManager) NewServerTransaction(msg *message.SipMsg) *ServerTransaction {

	tx := new(ServerTransaction)

	tx.TransManager = tm

    tx.InitFSM(msg)

	tx.BranchID = string(msg.Via[0].Branch)
	tx.Origin = msg
	return tx

}


