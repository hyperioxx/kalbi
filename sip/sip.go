package sip

import (
	"github.com/KalbiProject/Kalbi/sdp"
	"github.com/KalbiProject/Kalbi/sip/message"
	
)


func Parse(msg []byte) message.SipMsg {
	sipMsg := message.Parse(msg)
	sdp := sdp.Parse(msg)
	sipMsg.Sdp = sdp
	return sipMsg
}
