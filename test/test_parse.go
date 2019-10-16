package main

import (
	//"fmt"
		parser "GoSIP/sip"
	   )


func main(){
	var message string = `REGISTER sip:127.0.0.1;transport=UDP SIP/2.0\n\r
    Via: SIP/2.0/UDP 79.67.45.128:48189;branch=z9hG4bK-524287-1---3e839db9fc45b2cc;rport\n\r
	Max-forwards: 70\n\r
	Contact: <sip:43210@79.67.45.128:48189;rinstance=d3f929033197d10a;transport=UDP>\n\r
	To: "sdasdasd"<sip:43210@127.0.0.1;transport=UDP>\n\r
	From: "sdasdasd"<sip:43210@127.0.0.1;transport=UDP>;tag=2c8fbf43\n\r
	Call-id: _LMPeTigreTCSC3D0B32zw..\n\r
	Cseq: 8 REGISTER\n\r
	Expires: 60\n\r
	Allow: INVITE, ACK, CANCEL, BYE, NOTIFY, REFER, MESSAGE, OPTIONS, INFO, SUBSCRIBE\n\r
	User-agent: Z 5.2.28 rv2.8.115\n\r
	Allow-events: presence, kpml, talk\n\r
	Content-length: 0\n\r
	\n\r`

	//fmt.Println(message)

	x := parser.SipMessage{Raw:message}
	x.Parse(message)
}