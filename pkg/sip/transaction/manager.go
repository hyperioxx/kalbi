package transaction

import (
	//"fmt"
	"Kalbi/pkg/log"
	"Kalbi/pkg/sip/message"
	"Kalbi/pkg/transport"
	"github.com/looplab/fsm"
)



//NewTransactionManager returns a new TransactionManager
func NewTransactionManager() *TransactionManager {
	txmng := new(TransactionManager)
	txmng.tx = make(map[string]Transaction)
	return txmng
}





//TransactionManager handles SIP transactions
type TransactionManager struct {
	tx         map[string] Transaction
	RequestChannel    chan Transaction
	ResponseChannel   chan Transaction
	ListeningPoint    transport.ListeningPoint
}



// Start runs TransManager
func (tm *TransactionManager) Handle(request *message.SipMsg)  {
	
		if  request.Req.StatusCode != nil {
			log.Log.Info("Client transaction")

			tx, exists := tm.FindTransaction(string(request.Via[0].Branch))
			if(exists){
				log.Log.Info("Client Transaction aready exists")
				tm.ResponseChannel <- tx 
				
			}else{
				tx = tm.NewClientTransaction(request)
				tx.SetListeningPoint(tm.ListeningPoint)
			    tm.ResponseChannel <- tx  
			}
		} else if request.Req.Method != nil {
			log.Log.Info("Server transaction")
			tx, exists := tm.FindTransaction(string(request.Via[0].Branch))
			if(exists){
				log.Log.Info("Server Transaction aready exists")
				tm.RequestChannel <- tx 
				
			}else{
				tx = tm.NewServerTransaction(request)
				tx.SetListeningPoint(tm.ListeningPoint)
				tm.RequestChannel <- tx
			}
		}	
        
}


func (tm *TransactionManager) FindTransaction(branch string) (Transaction, bool) {
	tx , exists := tm.tx[branch]
	return tx , exists
}


func (tm *TransactionManager) DeleteTransaction(branch string){
	delete(tm.tx, branch)
}


func (tm *TransactionManager) NewClientTransaction(msg *message.SipMsg) *ClientTransaction {

	tx := new(ClientTransaction)

	tx.FSM = fsm.NewFSM("", fsm.Events{
		{Name: "Calling", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Trying", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Proceeding", Src: []string{"Calling"}, Dst: "Proceeding"},
		{Name: "Completed", Src: []string{"Calling", "Proceeding"}, Dst: "Completed"},
		{Name: "Terminated", Src: []string{"Calling", "Proceeding", "Completed"}, Dst: "Terminated"},
	}, fsm.Callbacks{})

	tx.BranchID = string(msg.Via[0].Branch)
	tx.Origin = msg

	tm.tx[string(msg.Via[0].Branch)] = tx

	return tx

}

func (tm *TransactionManager) NewServerTransaction(msg *message.SipMsg) *ServerTransaction {

	tx := new(ServerTransaction)

	tx.FSM = fsm.NewFSM("", fsm.Events{
		{Name: "Trying", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Proceeding", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Completed", Src: []string{"Calling", "Proceeding"}, Dst: "Completed"},
		{Name: "Terminated", Src: []string{"Calling", "Proceeding", "Completed"}, Dst: "Terminated"},
	}, fsm.Callbacks{})

	tx.BranchID = string(msg.Via[0].Branch)
	tx.Origin = msg

	tm.tx[string(msg.Via[0].Branch)] = tx

	return tx

}


