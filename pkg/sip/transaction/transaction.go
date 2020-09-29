package transaction

import (
	//"Kalbi/pkg/log"
	//"Kalbi/pkg/sdp"
	//"Kalbi/pkg/sip/message"
	//"Kalbi/pkg/transport"
	"github.com/looplab/fsm"
	"github.com/marv2097/siprocket"
)


type Transaction interface {
	GetBranchId() string
}



type ClientTransaction struct {
	ID          string
	BranchID    string
    Origin      string
	FSM         *fsm.FSM
	msg_history []*siprocket.SipMsg
}

func (ct *ClientTransaction) GetBranchId() string{
    return ct.BranchID
}

func (ct *ClientTransaction) SendRequest(msg *siprocket.SipMsg) {
	ct.msg_history = append(ct.msg_history, msg)
}



type ServerTransaction struct {
	ID          string
	BranchID    string
	FSM         *fsm.FSM
	msg_history []*siprocket.SipMsg

}

func (st *ServerTransaction) GetBranchId() string{
    return st.BranchID
}



func (st *ServerTransaction) Respond(response string){


}
