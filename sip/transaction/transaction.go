package transaction

import "fmt"
import "github.com/looplab/fsm"
import "Kalbi/sip/message"
import "Kalbi/log"
import "Kalbi/transport"
//import "github.com/davecgh/go-spew/spew"




type Transaction interface {
    Receive(msg *message.Request)
}


type ClientTransaction struct {
	ID string
	FSM *fsm.FSM
	msg_history[] *message.Request
}

func (ct *ClientTransaction) Receive(msg *message.Request){
	 fmt.Println(msg.Method)
	 ct.msg_history = append(ct.msg_history, msg)
}




type ServerTransaction struct {
	ID string
	FSM *fsm.FSM
	msg_history[] *message.Request
}

func (st *ServerTransaction) Receive(msg *message.Request){
	//spew.Dump(msg)
	st.msg_history = append(st.msg_history, msg)
	if msg.Method == "CANCEL" {
		response := message.NewResponse(200, msg)
		port := msg.Contact.Port
		transport.UdpSend(msg.Contact.Host, port, response)
	}else if st.FSM.Current() == "" {
		st.FSM.Event("Proceeding")
		msg.Headers.Set("Content-Length", "0")
	    //response := message.NewResponse(100, msg)
		port := msg.Contact.Port
		log.Log.Info("returning response to : " + msg.Contact.Host + ":" + port)
		//transport.UdpSend(msg.Contact.Host, port, response)
		response := message.NewResponse(200, msg)
		transport.UdpSend(msg.Contact.Host, port, response)

	}

}

func (st *ServerTransaction) Run(e *fsm.Event){
	
		log.Log.Info(e)
		//st.FSM.Event("open")
  
}








