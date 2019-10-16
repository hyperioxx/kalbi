package sip


type SipMessage struct {
	raw string 
	headers map[string][]SIPHeader

}