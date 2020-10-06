package transport

import (
	"Kalbi/pkg/sip/message"
	
)

type ListeningPoint interface {
	Read() *message.SipMsg
	Build(string, int)
}
