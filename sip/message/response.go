package message

import "fmt"
var TRYING_100 string = "100 Trying"
var RINGING_180 string  = "180 Ringing"
var OK_200 string = "200 OK"



func NewResponse(code int ,request *Request) string {
	var response string

	switch code {
	case 100:
	  response += "SIP/2.0 " + TRYING_100 + "\r\n"
	case 200:
	  response += "SIP/2.0 " + OK_200 + "\r\n"
	}

	response = buildResponse(response, request)

    return response
}






func buildResponse(response string ,request *Request) string{

	for k,v := range request.Headers{
        response +=  k + " : " + v[0] +  "\r\n"
	}
	response +=  "\r\n\r\n"
	response += request.SDP.String()
	fmt.Println(response)
	return response
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

