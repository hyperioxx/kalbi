package message

import (
	"github.com/marv2097/siprocket"
	//"fmt"
)
var TRYING_100 string = "100 Trying"
var RINGING_180 string  = "180 Ringing"
var OK_200 string = "200 OK"



func NewResponse(code int ,request *siprocket.SipMsg) string {
	var response string

	switch code {
	case 100:
	  response = response + "SIP/2.0 " + TRYING_100 + "\r\n"
	case 200:
	  response = response + "SIP/2.0 " + OK_200 + "\r\n"
	}

	msg := buildResponse(response, request)

    return msg
}






func buildResponse(response string ,request *siprocket.SipMsg) string {
	response_message := ""
	response_message += response
	for _, header := range request.Via{
	    response_message += "Via: " + string(header.Src) + "\r\n"
	}
	response_message += "From: " + string(request.From.Src) + "\r\n"
	response_message += "To: " + string(request.To.Src) + "\r\n"
	response_message += "Cseq: " + string(request.Cseq.Src) + "\r\n"
	response_message += "Call-ID: " + string(request.CallId.Value) + "\r\n"
	response_message += "Server: Kalbi Server 0.1" + "\r\n"
	response_message += "Contact: " + string(request.Contact.Src) + "\r\n"

	response_message += "\r\n\r\n"
	
	return response_message

}




//Request sip request struct
type Response struct {
	ResponseCode     string
	RequestURI string
	Proto      string
	Headers    map[string]string
	Via        *URI
	Contact    *URI
	To         *URI
	From       *URI
	Branch     string
	Cseq       string

}

