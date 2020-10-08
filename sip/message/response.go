package message

import (
	"github.com/KalbiProject/Kalbi/sip/status"
	"strings"
)

func NewResponse(response int, to string, from string) *SipMsg {
	//TODO: need more elegant way to create responses
	to_ := strings.Split(to, "@")
	from_ := strings.Split(from, "@")

	r := new(SipMsg)
	r.Req = *new(SipReq)
	r.To = *new(SipTo)
	r.From = *new(SipFrom)

	r.Req.SetStatusCode(response)
	r.Req.SetStatusDesc(status.StatusText(response))
	r.Req.SetUriType("sip")
	r.Req.SetUser(to_[0])
	r.Req.SetHost(to_[1])

	r.To.SetUriType("sip")
	r.To.SetUser(to_[0])
	r.To.SetHost(to_[1])

	r.From.SetUriType("sip")
	r.From.SetUser(from_[0])
	r.From.SetHost(from_[1])

	r.ContLen = *new(SipVal)
	return r
}
