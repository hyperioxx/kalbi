package interfaces

import (
	"github.com/KalbiProject/Kalbi/sip/message"
)

type SipListener interface {
	HandleRequests(SipEventObject)
	HandleResponses(SipEventObject)
}

type Transaction interface {
	GetBranchID() string
	GetOrigin() *message.SipMsg
	SetListeningPoint(ListeningPoint)
	Send(*message.SipMsg, string, string)
	Receive(*message.SipMsg)
	GetLastMessage() *message.SipMsg
	GetServerTransactionID() string
	SetLastMessage(*message.SipMsg)
	GetListeningPoint() ListeningPoint
}

type SipEventObject interface {
	GetSipMessage() *message.SipMsg
	SetSipMessage(*message.SipMsg)
	GetTransaction() Transaction
	SetTransaction(Transaction)
}

type ListeningPoint interface {
	Read() SipEventObject
	Build(string, int)
	Start()
	SetTransportChannel(chan SipEventObject)
	Send(string, string, string) error
}
