package parser

import(
		"github.com/marv2097/siprocket"
        
        )



func Read(msg []byte) *siprocket.SipMsg {

	sip := siprocket.Parse(msg)

    return &sip
}
