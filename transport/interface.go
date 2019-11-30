package transport

import "Kalbi/sip/message"

type ListeningPoint interface {
    Read() *message.Request
    Build(string, int)
}