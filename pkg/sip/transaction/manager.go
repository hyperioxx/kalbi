package transaction

import (
	"github.com/looplab/fsm"
	"github.com/marv2097/siprocket"
)

//TransactionManager handles SIP transactions
type TransactionManager struct {
	ctx   map[string]*ClientTransaction
	stx   map[string]*ServerTransaction
	input chan siprocket.SipMsg
}

func (tm *TransactionManager) Init() {
	tm.ctx = make(map[string]*ClientTransaction)
	tm.stx = make(map[string]*ServerTransaction)
}

//SetChannels setup input & output channels for
func (tm *TransactionManager) SetChannel(input chan siprocket.SipMsg) {
	tm.input = input
}

// Start runs TransManager
func (tm *TransactionManager) Start() {
	for {
		var trans *ServerTransaction
		request := <-tm.input

		trans = tm.NewServerTransaction(&request)

		trans.Receive(&request)

	}
}


func (tm *TransactionManager) FindClientTransaction(branch string) *ClientTransaction {
	return tm.ctx[branch]
}


func (tm *TransactionManager) FindServerTransaction(branch string) *ServerTransaction {
	if val, ok := tm.stx[branch]; ok {
		return val
	}
	return nil
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
	
	//tm.stx[msg.Branch] = tx

	return tx

}

func NewManager(input chan siprocket.SipMsg) *Manager {
	transManager := new(Manager)
	transManager.SetChannel(input)
	transManager.Init()
	return transManager
}
