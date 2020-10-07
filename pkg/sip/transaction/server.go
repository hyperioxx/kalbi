package transaction

import (
	//"fmt"
	"github.com/KalbiProject/Kalbi/pkg/sip/method"
	"github.com/KalbiProject/Kalbi/pkg/sip/message"
	"github.com/KalbiProject/Kalbi/pkg/log"
	"github.com/KalbiProject/Kalbi/pkg/transport"
	"github.com/looplab/fsm"
	
)

const (
	server_input_request = "server_input_request"
	server_input_ack = "server_input_ack"
	server_input_user_1xx = "server_input_user_1xx"
	server_input_user_2xx = "server_input_user_2xx"
	server_input_user_300_plus = "server_input_user_300_plus"
	server_input_timer_g = "server_input_timer_g"
	server_input_timer_h = "server_input_timer_h"
	server_input_timer_i = "server_input_timer_i"
	server_input_transport_err = "server_input_transport_err"
	server_input_delete = "server_input_delete"
)




type ServerTransaction struct {
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

}


func (st *ServerTransaction) InitFSM(msg *message.SipMsg){

	switch string(msg.Req.Method) {
	case method.INVITE:
		st.FSM = fsm.NewFSM("", fsm.Events{
			{Name: server_input_request, Src: []string{""}, Dst: "Proceeding"},
			{Name: server_input_user_1xx, Src: []string{"Proceeding"}, Dst: "Proceeding"},
			{Name: server_input_user_300_plus , Src: []string{"Proceeding"}, Dst: "Completed"},
			{Name: server_input_ack, Src: []string{"Completed"}, Dst: "Confirmed"},
			{Name: server_input_user_2xx, Src: []string{"Proceeding"}, Dst: "Terminated"},
		}, fsm.Callbacks{server_input_user_1xx : st.actRespond,
						 server_input_transport_err : st.actTransErr,
						 server_input_user_2xx: st.actRespondDelete, 
						 server_input_user_300_plus: st.actRespond })
	default:
		st.FSM = fsm.NewFSM("", fsm.Events{
			{Name: server_input_request, Src: []string{""}, Dst: "Proceeding"},
			{Name: server_input_user_1xx, Src: []string{"Proceeding"}, Dst: "Proceeding"},
			{Name: server_input_user_300_plus , Src: []string{"Proceeding"}, Dst: "Completed"},
			{Name: server_input_ack, Src: []string{"Completed"}, Dst: "Confirmed"},
			{Name: server_input_user_2xx, Src: []string{ "Proceeding"}, Dst: "Terminated"},
		}, fsm.Callbacks{server_input_user_1xx : st.actRespond,
			server_input_transport_err : st.actTransErr,
			server_input_user_2xx: st.actRespondDelete, 
			server_input_user_300_plus: st.actRespond})
}
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

func (st *ServerTransaction) Receive(msg *message.SipMsg){
	  st.LastMessage = msg
	  log.Log.Info("Message Received for transactionId "+ st.BranchID +": \n" + string(msg.Src))
	  log.Log.Info(message.MessageDetails(msg))
	  if msg.Req.Method != nil || string(msg.Req.Method) != method.ACK {
		  st.FSM.Event(server_input_request)
	  }

}

func (st *ServerTransaction)Respond(msg *message.SipMsg){
	//TODO: this will change due to issue https://github.com/KalbiProject/Kalbi/issues/20
	log.Log.Info("Message Sent for transactionId "+ st.BranchID +": \n" + message.MessageDetails(msg))
	if msg.GetStatusCode() < 200 { 
		st.FSM.Event(server_input_user_1xx)
	}else if msg.GetStatusCode() < 300 {
		st.FSM.Event(server_input_user_2xx)
	}else {
	    st.FSM.Event(server_input_user_300_plus)
	}

}


func (st *ServerTransaction) Send(msg *message.SipMsg, host string, port string) {
	st.LastMessage = msg
	st.Host = host
	st.Port = port
	st.Respond(msg)
}


func (st *ServerTransaction) actRespond(event *fsm.Event) {
	err := st.ListeningPoint.Send(st.Host, st.Port, st.LastMessage.Export())
	if err != nil {
		st.FSM.Event(server_input_transport_err)
	}
    
}

func (st *ServerTransaction) actRespondDelete(event *fsm.Event){
	err := st.ListeningPoint.Send(st.Host, st.Port, st.LastMessage.Export())
	if err != nil {
		st.FSM.Event(server_input_transport_err)
	}
	st.TransManager.DeleteTransaction(st.BranchID)
}

func (st *ServerTransaction) actTransErr(event *fsm.Event) {
	log.Log.Error("Transport error for transactionID : "+ st.BranchID)
	st.FSM.Event(server_input_delete)
}

func (st *ServerTransaction) actDelete(event *fsm.Event) {
    st.TransManager.DeleteTransaction(st.BranchID)
}


