package transaction

import (
	"Kalbi/pkg/sip/message"
	"Kalbi/pkg/transport"
	"github.com/looplab/fsm"
	
)


type Transaction interface {
	GetBranchId()       string
	GetOrigin()         *message.SipMsg
	SetListeningPoint(transport.ListeningPoint)
	Send(*message.SipMsg, string, string)  
}



type ClientTransaction struct {
	ID               string
	BranchID         string
    Origin           *message.SipMsg
	FSM              *fsm.FSM
	ListeningPoint   transport.ListeningPoint

}

func (ct *ClientTransaction) SetListeningPoint(lp transport.ListeningPoint){
    ct.ListeningPoint = lp
}

func (ct *ClientTransaction) GetBranchId() string{
    return ct.BranchID
}

func (ct *ClientTransaction) GetOrigin() *message.SipMsg{
    return ct.Origin
}

func (ct *ClientTransaction) Send(msg *message.SipMsg, host string, port string) {
	ct.ListeningPoint.Send(host, port, msg.Export())
}



type ServerTransaction struct {
	ID               string
	BranchID         string
	Origin           *message.SipMsg
	FSM              *fsm.FSM
	msg_history      []*message.SipMsg
	ListeningPoint   transport.ListeningPoint

}

func (st *ServerTransaction) SetListeningPoint(lp transport.ListeningPoint){
    st.ListeningPoint = lp
}

func (st *ServerTransaction) GetBranchId() string{
    return st.BranchID
}

func (st *ServerTransaction) GetOrigin() *message.SipMsg{
    return st.Origin
}

func (st *ServerTransaction) Send(msg *message.SipMsg, host string, port string) {
	st.ListeningPoint.Send(host, port, msg.Export())
}