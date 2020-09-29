package events

import "github.com/marv2097/siprocket"
import "Kalbi/pkg/sip/transaction"

func NewEvent(msg siprocket.SipMsg, tx transaction.Transaction) *Event{
	event := new(Event)
	event.Message = &msg
	event.Transaction = tx
	if string(msg.Req.StatusCode) != "" {
		event.Type = "Client"
        
	}else if string(msg.Req.Method) != "" {
		event.Type = "Server"
	}	
	return event
}


type Event struct {
	Type string
	Message *siprocket.SipMsg
	Transaction transaction.Transaction
}


