package sip

import (
	"github.com/KalbiProject/kalbi/sip/message"
)

//Parse parses SIP message
func Parse(msg []byte) message.SipMsg {
	sipMsg := message.Parse(msg)
	return sipMsg
}
