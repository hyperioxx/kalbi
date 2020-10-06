package events

import "Kalbi/pkg/sip/transaction"
import "Kalbi/pkg/sip/message"

func NewEvent(msg message.SipMsg, tx transaction.Transaction) *Event{
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
	Message *message.SipMsg
	Transaction transaction.Transaction
}


