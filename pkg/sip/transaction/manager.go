package transaction

import (
	//"fmt"
	"Kalbi/pkg/log"
	"github.com/looplab/fsm"
	"github.com/marv2097/siprocket"
)


//NewTransactionManager returns a new TransactionManager
func NewTransactionManager() *TransactionManager {
	txmng := new(TransactionManager)
	txmng.tx = make(map[string]Transaction)
	return txmng
}





//TransactionManager handles SIP transactions
type TransactionManager struct {
	tx   map[string] Transaction
	input chan siprocket.SipMsg
	output chan siprocket.SipMsg
}





// Start runs TransManager
func (tm *TransactionManager) Handle(request *siprocket.SipMsg)  {
	
		if string(request.Req.StatusCode) != "" {
			log.Log.Info("Client transaction")
			tm.NewClientTransaction(request)

		} else if string(request.Req.Method) != "" {
			log.Log.Info("Server transaction")
			tm.NewServerTransaction(request)
		}	
        
}


func (tm *TransactionManager) FindTransaction(branch string) Transaction {
	return tm.tx[branch]
}


func (tm *TransactionManager) DeleteTransaction(branch string){
	delete(tm.tx, branch)
}


func (tm *TransactionManager) NewClientTransaction(msg *siprocket.SipMsg) *ClientTransaction {

	tx := new(ClientTransaction)

	tx.FSM = fsm.NewFSM("", fsm.Events{
		{Name: "Calling", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Trying", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Proceeding", Src: []string{"Calling"}, Dst: "Proceeding"},
		{Name: "Completed", Src: []string{"Calling", "Proceeding"}, Dst: "Completed"},
		{Name: "Terminated", Src: []string{"Calling", "Proceeding", "Completed"}, Dst: "Terminated"},
	}, fsm.Callbacks{})

	tx.BranchID = string(msg.Via[0].Branch)

	tm.tx[string(msg.Via[0].Branch)] = tx

	return tx

}

func (tm *TransactionManager) NewServerTransaction(msg *siprocket.SipMsg) *ServerTransaction {

	tx := new(ServerTransaction)

	tx.FSM = fsm.NewFSM("", fsm.Events{
		{Name: "Trying", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Proceeding", Src: []string{""}, Dst: "Proceeding"},
		{Name: "Completed", Src: []string{"Calling", "Proceeding"}, Dst: "Completed"},
		{Name: "Terminated", Src: []string{"Calling", "Proceeding", "Completed"}, Dst: "Terminated"},
	}, fsm.Callbacks{})

	tx.BranchID = string(msg.Via[0].Branch)

	tm.tx[string(msg.Via[0].Branch)] = tx

	return tx

}


