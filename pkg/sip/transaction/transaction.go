package transaction

import (
	"github.com/KalbiProject/Kalbi/pkg/sip/message"
	"github.com/KalbiProject/Kalbi/pkg/transport"
	
)


type Transaction interface {
	GetBranchId()       string
	GetOrigin()         *message.SipMsg
	SetListeningPoint(transport.ListeningPoint)
	Send(*message.SipMsg, string, string)
	Receive(*message.SipMsg)  
}







