package transaction

import (
	"Kalbi/pkg/sip/message"
	"Kalbi/pkg/transport"
	
)


type Transaction interface {
	GetBranchId()       string
	GetOrigin()         *message.SipMsg
	SetListeningPoint(transport.ListeningPoint)
	Send(*message.SipMsg, string, string)
	Receive(*message.SipMsg)  
}







