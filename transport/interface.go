package transport

import (
	"github.com/KalbiProject/Kalbi/sip/message"
)

type ListeningPoint interface {
	Read() *message.SipMsg
	Build(string, int)
	Start()
	Send(string, string, string) error
}
