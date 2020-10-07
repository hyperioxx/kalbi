package transport

import (
	"github.com/KalbiProject/Kalbi/pkg/sip/message"
	
)

type ListeningPoint interface {
	Read() *message.SipMsg
	Build(string, int)
	Send( string, string, string) error
}
