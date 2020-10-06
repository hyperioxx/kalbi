package transaction

import (
	//"Kalbi/pkg/log"
	//"Kalbi/pkg/sdp"
	"Kalbi/pkg/sip/message"
	//"Kalbi/pkg/transport"
	"github.com/looplab/fsm"
	
)


type Transaction interface {
	GetBranchId() string
}



type ClientTransaction struct {
	ID          string
	BranchID    string
    Origin      string
	FSM         *fsm.FSM
	msg_history []*message.SipMsg
}

func (ct *ClientTransaction) GetBranchId() string{
    return ct.BranchID
}

func (ct *ClientTransaction) SendRequest(msg *message.SipMsg) {
	ct.msg_history = append(ct.msg_history, msg)
}



type ServerTransaction struct {
	ID          string
	BranchID    string
	FSM         *fsm.FSM
	msg_history []*message.SipMsg

}

func (st *ServerTransaction) GetBranchId() string{
    return st.BranchID
}



func (st *ServerTransaction) Respond(response string){

}
