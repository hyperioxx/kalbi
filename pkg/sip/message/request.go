package message

import "net/textproto"
import "github.com/pixelbender/go-sdp/sdp"

//Request sip request struct
type Request struct {
	Method     string
	RequestURI string
	Proto      string
	Headers    textproto.MIMEHeader
	Via        *URI
	Contact    *URI
	To         *URI
	From       *URI
	Branch     string
	Cseq       string
	CallID     string
	SDP        *sdp.Session
}

func (r *Request) GetHeader(header string) string {
	if val, ok := r.Headers[header]; ok {
		return val[0]
	}
	return ""
}

func (r *Request) SetHeader(header string, value string) string {
	if val, ok := r.Headers[header]; ok {
		return val[0]
	}
	return ""
}
