package transaction

/* 17.1 Client Transaction

   The client transaction provides its functionality through the
   maintenance of a state machine.

   The TU communicates with the client transaction through a simple
   interface.  When the TU wishes to initiate a new transaction, it
   creates a client transaction and passes it the SIP request to send
   and an IP address, port, and transport to which to send it.  The
   client transaction begins execution of its state machine.  Valid
   responses are passed up to the TU from the client transaction.

   There are two types of client transaction state machines, depending
   on the method of the request passed by the TU.  One handles client
   transactions for INVITE requests.  This type of machine is referred
   to as an INVITE client transaction.  Another type handles client
   transactions for all requests except INVITE and ACK.  This is
   referred to as a non-INVITE client transaction.  There is no client
   transaction for ACK.  If the TU wishes to send an ACK, it passes one
   directly to the transport layer for transmission.

   The INVITE transaction is different from those of other methods
   because of its extended duration.  Normally, human input is required
   in order to respond to an INVITE.  The long delays expected for
   sending a response argue for a three-way handshake.  On the other
   hand, requests of other methods are expected to complete rapidly.
   Because of the non-INVITE transaction's reliance on a two-way
   handshake, TUs SHOULD respond immediately to non-INVITE requests. */

import (
	"github.com/KalbiProject/Kalbi/log"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/transport"
	"github.com/looplab/fsm"
	"time"
)

const (
	client_input_request       = "client_input_request"
	client_input_1xx           = "client_input_1xx"
	client_input_2xx           = "client_input_2xx"
	client_input_300_plus      = "client_input_300_plus"
	client_input_timer_a       = "client_input_timer_a"
	client_input_timer_b       = "client_input_timer_b"
	client_input_timer_d       = "client_input_timer_d"
	client_input_transport_err = "client_input_transport_err"
	client_input_delete        = "client_input_transport_err"
)

type ClientTransaction struct {
	ID             string
	BranchID       string
	TransManager   *TransactionManager
	Origin         *message.SipMsg
	FSM            *fsm.FSM
	msg_history    []*message.SipMsg
	ListeningPoint transport.ListeningPoint
	Host           string
	Port           string
	LastMessage    *message.SipMsg
	timer_a_time   time.Duration
	timer_a        *time.Timer
	timer_b        *time.Timer
	timer_d_time   time.Duration
	timer_d        *time.Timer
}

func (ct *ClientTransaction) InitFSM(msg *message.SipMsg) {

	switch string(msg.Req.Method) {
	case method.INVITE:
		ct.FSM = fsm.NewFSM("", fsm.Events{
			{Name: client_input_request, Src: []string{""}, Dst: "Calling"},
			{Name: client_input_1xx, Src: []string{"Calling"}, Dst: "Proceeding"},
			{Name: client_input_300_plus, Src: []string{"Proceeding"}, Dst: "Completed"},
			{Name: client_input_2xx, Src: []string{"Proceeding"}, Dst: "Terminated"},
			{Name: client_input_transport_err, Src: []string{"Calling", "Proceeding", "Completed"}, Dst: "Terminated"},
		}, fsm.Callbacks{client_input_2xx: ct.actDelete,
			client_input_300_plus: ct.act300,
			client_input_timer_a:  ct.actResend,
			client_input_timer_b:  ct.actTransErr,
		})

	default:
		ct.FSM = fsm.NewFSM("", fsm.Events{
			{Name: client_input_request, Src: []string{""}, Dst: "Calling"},
			{Name: client_input_1xx, Src: []string{"Calling"}, Dst: "Proceeding"},
			{Name: client_input_300_plus, Src: []string{"Proceeding"}, Dst: "Completed"},
			{Name: client_input_2xx, Src: []string{"Proceeding"}, Dst: "Terminated"},
		}, fsm.Callbacks{})
	}
}

func (ct *ClientTransaction) SetListeningPoint(lp transport.ListeningPoint) {
	ct.ListeningPoint = lp
}

func (ct *ClientTransaction) GetBranchId() string {
	return ct.BranchID
}

func (ct *ClientTransaction) GetOrigin() *message.SipMsg {
	return ct.Origin
}

func (ct *ClientTransaction) Receive(msg *message.SipMsg) {

	if msg.GetStatusCode() < 200 {
		ct.FSM.Event(server_input_user_1xx)
	} else if msg.GetStatusCode() < 300 {
		ct.FSM.Event(server_input_user_2xx)
	} else {
		ct.FSM.Event(server_input_user_300_plus)
	}

}

func (ct *ClientTransaction) actSend(event *fsm.Event) {
	err := ct.ListeningPoint.Send(ct.Host, ct.Port, ct.Origin.Export())
	if err != nil {
		ct.FSM.Event(client_input_transport_err)
	}
}

func (ct *ClientTransaction) act300(event *fsm.Event) {
	log.Log.Debug("Client transaction %p, act_300", ct)
	ct.timer_d = time.AfterFunc(ct.timer_d_time, func() {
		ct.FSM.Event(client_input_timer_d)
	})

}

func (ct *ClientTransaction) actTransErr(event *fsm.Event) {
	log.Log.Error("Transport error for transactionID : " + ct.BranchID)
	ct.FSM.Event(client_input_delete)
}

func (ct *ClientTransaction) actDelete(event *fsm.Event) {
	ct.TransManager.DeleteTransaction(string(ct.Origin.Via[0].Branch))
}

func (ct *ClientTransaction) actResend(event *fsm.Event) {
	log.Log.Debug("Client transaction %p, act_resend", ct)
	ct.timer_a_time *= 2
	ct.timer_a.Reset(ct.timer_a_time)
	ct.Resend()
}

func (ct *ClientTransaction) Resend() {
	err := ct.ListeningPoint.Send(ct.Host, ct.Port, ct.Origin.Export())
	if err != nil {
		ct.FSM.Event(client_input_transport_err)
	}
}

func (ct *ClientTransaction) StatelessSend(msg *message.SipMsg, host string, port string) {
	err := ct.ListeningPoint.Send(ct.Host, ct.Port, ct.Origin.Export())

	if err != nil {
		log.Log.Error("Transport error for transactionID : " + ct.BranchID)
	}

}

func (ct *ClientTransaction) Send(msg *message.SipMsg, host string, port string) {
	defer ct.FSM.Event(server_input_request)
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

	err := ct.ListeningPoint.Send(ct.Host, ct.Port, ct.Origin.Export())
	if err != nil {
		ct.FSM.Event(server_input_transport_err)
	}

}
