package sip


type SipMessage struct {
	Raw string 
	Headers map[string][]SIPHeader
}