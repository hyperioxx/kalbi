package transaction

import "fmt"
import "Kalbi/sip/message"

//TransactionManager handles SIP transactions
type Manager struct {
	txs map[string] *Transaction
	txc map[string] *Transaction
	input chan *message.Request
}

func (tm *Manager) init(){
	tm.txs = make(map[string] *Transaction)
	tm.txc = make(map[string] *Transaction)
}

//SetChannels setup input & output channels for 
func (tm *Manager) SetChannel(input chan *message.Request){
    tm.input = input
}

// Start runs TransManager
func (tm *Manager) Start(){
	for{
		request := <- tm.input
        fmt.Println(request)
	}


}


func NewManager(input chan *message.Request) *Manager{
	transManager := new(Manager)
	transManager.SetChannel(input)
	return transManager
}





