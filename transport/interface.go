package transport

import (
	"github.com/KalbiProject/Kalbi/sip/message"
)

type ListeningPoint interface {
	Read() *message.SipMsg
	Build(string, int)
	Start()
	SetTransportChannel(chan *message.SipMsg)
	Send(string, string, string) error
}
