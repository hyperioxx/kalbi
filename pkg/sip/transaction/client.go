package transaction

import (
	"time"
	"github.com/KalbiProject/Kalbi/pkg/sip/message"
	"github.com/KalbiProject/Kalbi/pkg/transport"
	"github.com/looplab/fsm"
	"github.com/KalbiProject/Kalbi/pkg/sip/method"
	
)


const (
	client_input_1xx = "client_input_1xx"
	client_input_2xx = "client_input_2xx"
	client_input_300_plus = "client_input_300_plus"
	client_input_timer_a = "client_input_timer_a"
	client_input_timer_b = "client_input_timer_b"
	client_input_timer_d = "client_input_timer_d"
	client_input_transport_err = "client_input_transport_err"
	client_input_delete = "client_input_transport_err"
)

type ClientTransaction struct {
	ID               string
	BranchID         string
	TransManager     *TransactionManager
	Origin           *message.SipMsg
	FSM              *fsm.FSM
	msg_history      []*message.SipMsg
	ListeningPoint   transport.ListeningPoint
	Host             string
	Port             string
	LastMessage      *message.SipMsg
	timer_a_time     time.Duration
	timer_a          *time.Timer
	timer_b          *time.Timer
	timer_d_time     time.Duration 
	timer_d          time.Timer
}

func (ct *ClientTransaction) InitFSM(msg *message.SipMsg){

	switch string(msg.Req.Method) {
	case method.INVITE:
		ct.FSM = fsm.NewFSM("", fsm.Events{
			{Name: server_input_request, Src: []string{""}, Dst: "Calling"},
			{Name: server_input_user_1xx, Src: []string{"Calling"}, Dst: "Proceeding"},
			{Name: server_input_user_300_plus , Src: []string{"Proceeding"}, Dst: "Completed"},
			{Name: server_input_ack, Src: []string{"Completed"}, Dst: "Confirmed"},
			{Name: server_input_user_2xx, Src: []string{"Proceeding"}, Dst: "Terminated"},
		}, fsm.Callbacks{})
	
	default:
		ct.FSM = fsm.NewFSM("", fsm.Events{
			{Name: server_input_request, Src: []string{""}, Dst: "Proceeding"},
			{Name: server_input_user_1xx, Src: []string{"Proceeding"}, Dst: "Proceeding"},
			{Name: server_input_user_300_plus , Src: []string{"Proceeding"}, Dst: "Completed"},
			{Name: server_input_ack, Src: []string{"Completed"}, Dst: "Confirmed"},
			{Name: server_input_user_2xx, Src: []string{ "Proceeding"}, Dst: "Terminated"},
		}, fsm.Callbacks{})
}
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

func (ct *ClientTransaction) Receive(msg *message.SipMsg){

	if msg.GetStatusCode() < 200 { 
		ct.FSM.Event(server_input_user_1xx)
	}else if msg.GetStatusCode() < 300 {
		ct.FSM.Event(server_input_user_2xx)
	}else {
	    ct.FSM.Event(server_input_user_300_plus)
	}
	  
}


func (ct *ClientTransaction) actSend(msg *message.SipMsg){
	err := ct.ListeningPoint.Send(ct.Host, ct.Port, ct.Origin.Export())
	if err != nil {
		ct.FSM.Event(server_input_transport_err)
	}
}

func (ct *ClientTransaction) resend() {
	err := ct.ListeningPoint.Send(ct.Host, ct.Port, ct.Origin.Export())
	if err != nil {
		ct.FSM.Event(client_input_transport_err)
	}
}




func (ct *ClientTransaction) Send(msg *message.SipMsg, host string, port string) {
	ct.Origin = msg
	ct.Host = host
	ct.Port = port
	ct.timer_a_time = T1

	//Retransmition timer
	ct.timer_a = time.AfterFunc(ct.timer_a_time, func() {
	   ct.FSM.Event(client_input_timer_a)
	})

	//timeout timer
	ct.timer_b = time.AfterFunc(64*T1, func() {
		ct.FSM.Event(client_input_timer_b)
	})

	ct.actSend(msg)
}