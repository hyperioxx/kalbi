package transaction

import (
	"Kalbi/internal/kalbi/log"
	"Kalbi/internal/kalbi/sdp"
	"Kalbi/internal/kalbi/sip/message"
	"Kalbi/internal/kalbi/transport"
	"github.com/looplab/fsm"
	"github.com/marv2097/siprocket"
	
)

//import "github.com/davecgh/go-spew/spew"

type Transaction interface {
	GetBranchId()
	GetDialog()
	GetRequest()
	GetState()
	
}



type ClientTransaction struct {
	ID          string
	FSM         *fsm.FSM
	msg_history []*siprocket.SipMsg
}



func (ct *ClientTransaction) SendRequest(msg *siprocket.SipMsg) {
	ct.msg_history = append(ct.msg_history, msg)
}









type ServerTransaction struct {
	ID          string
	FSM         *fsm.FSM
	msg_history []*siprocket.SipMsg
}

func (st *ServerTransaction) Receive(msg *siprocket.SipMsg) {
	st.msg_history = append(st.msg_history, msg)
	if string(msg.Req.Method) == "ACK" {
		return
	}
	if string(msg.Req.Method) == "CANCEL" || string(msg.Req.Method) == "BYE" {
		response := message.NewResponse(200, msg)
		port := string(msg.Contact.Port)
		transport.UdpSend(string(msg.Contact.Host), string(port), response)
	} else if st.FSM.Current() == "" {
		st.FSM.Event("Proceeding")
		response := message.NewResponse(100, msg)
		port := msg.Contact.Port
		log.Log.Info("returning response to : " + string(msg.Contact.Host) + ":" + string(port))
		transport.UdpSend(string(msg.Contact.Host), string(port), response)
		sdp.HandleSdp(msg.Sdp)
		response = message.NewResponse(200, msg)
		transport.UdpSend(string(msg.Contact.Host), string(port), response)

	}

}

func (st *ServerTransaction) Respond(response string){


}
