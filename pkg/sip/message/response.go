package message

import (
	"github.com/marv2097/siprocket"
	"Kalbi/pkg/sip/status"
)



func NewResponse(code int) Response {
	var response Response

	switch code {
	case 100:
		response.ResponseCode =  "SIP/2.0 " + status.TRYING_100 + "\r\n"
	case 200:
		response.ResponseCode =  "SIP/2.0 "+ status.OK_200 + "\r\n"
	}

	return response
}

func buildResponse(response string) Response {

	return Response

}

//Request sip request struct
type Response struct {
	ResponseCode string
	RequestURI   string
	Proto        string
	Headers      map[string]string
	Via          []*URI
	Contact      []*URI
	To           *URI
	From         *URI
	Branch       string
	Cseq         string
}


func (r *Response) buildResponse(response string) string {
	response_message := ""
	response_message += response
	for _, header := range r.Via {
		response_message += "Via: " + string(header.Scheme) + "\r\n"
	}
	response_message += "From: " + string(request.From.Src) + "\r\n"
	response_message += "To: " + string(request.To.Src) + "\r\n"
	response_message += "Cseq: " + string(request.Cseq.Src) + "\r\n"
	response_message += "Call-ID: " + string(request.CallId.Value) + "\r\n"
	response_message += "Contact: " + string(request.Contact.Src) + "\r\n"

	response_message += "\r\n\r\n"

	return response_message

}

