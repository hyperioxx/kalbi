package transaction

import (
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/transport"
)

type Transaction interface {
	GetBranchId() string
	GetOrigin() *message.SipMsg
	SetListeningPoint(transport.ListeningPoint)
	Send(*message.SipMsg, string, string)
	Receive(*message.SipMsg)
}
