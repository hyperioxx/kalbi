package message

import (
	"strings"
)

func NewRequest(request string, to string, from string) *SipMsg {
	//TODO: need more elegant way to create messages
	to_ := strings.Split(to, "@")
	from_ := strings.Split(from, "@")

	r := new(SipMsg)
	via := new(SipVia)
	r.Via = make([]SipVia, 1)
	r.Via = append(r.Via, *via)
	r.Req = *new(SipReq)
	r.To = *new(SipTo)
	r.From = *new(SipFrom)
	r.Cseq = *new(SipCseq)
	r.Contact = *new(SipContact)

	via.SetHost(from_[0])
	via.SetPort(from_[1])
	via.SetBranch(GenerateBranchId())

	r.Req.SetMethod(request)
	r.Req.SetUriType("sip")
	r.Req.SetUser(to_[0])
	r.Req.SetHost(to_[1])

	r.To.SetUriType("sip")
	r.To.SetUser(to_[0])
	r.To.SetHost(to_[1])

	r.From.SetUriType("sip")
	r.From.SetUser(from_[0])
	r.From.SetHost(from_[1])

	r.Cseq.SetID("1")
	r.Cseq.SetMethod(request)
	r.ContLen.SetValue("0")
	return r
}
