package message

import (
	"github.com/KalbiProject/Kalbi/sip/status"
	"strings"
)

//NewResponse creates new SIP response
func NewResponse(response int, to string, from string) *SipMsg {
	//TODO: need more elegant way to create responses
	toArr := strings.Split(to, "@")
	fromArr := strings.Split(from, "@")
	toUser, toDomain := toArr[0], toArr[1]
	fromUser, fromDomain := fromArr[0], fromArr[1]

	r := new(SipMsg)
	r.Req = *new(SipReq)
	r.To = *new(SipTo)
	r.From = *new(SipFrom)

	r.Req.SetStatusCode(response)
	r.Req.SetStatusDesc(status.StatusText(response))
	r.Req.SetUriType("sip")
	r.Req.SetUser(toUser)
	r.Req.SetHost(toDomain)

	r.To.SetUriType("sip")
	r.To.SetUser(toUser)
	r.To.SetHost(toDomain)

	r.From.SetUriType("sip")
	r.From.SetUser(fromUser)
	r.From.SetHost(fromDomain)

	r.ContLen = *new(SipVal)
	return r
}
