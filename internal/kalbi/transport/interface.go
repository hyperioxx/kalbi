package transport

import (
	"github.com/marv2097/siprocket"
)

type ListeningPoint interface {
	Read() *siprocket.SipMsg
	Build(string, int)
}
