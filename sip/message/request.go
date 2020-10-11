package message

import (
	"strings"
)

//NewRequest creates new SIP request
func NewRequest(request string, to string, from string) *SipMsg {
	//TODO: need more elegant way to create messages
	toArr := strings.Split(to, "@")
	fromArr := strings.Split(from, "@")
	toUser, toDomain := toArr[0], toArr[1]
	fromUser, fromDomain := fromArr[0], fromArr[1]

	r := new(SipMsg)
	via := new(SipVia)
	r.Via = make([]SipVia, 1)
	r.Via = append(r.Via, *via)
	r.Req = *new(SipReq)
	r.To = *new(SipTo)
	r.From = *new(SipFrom)
	r.Cseq = *new(SipCseq)
	r.Contact = *new(SipContact)

	via.SetHost(fromUser)
	via.SetPort(fromDomain)
	via.SetBranch(GenerateBranchId())

	r.Req.SetMethod(request)
	r.Req.SetUriType("sip")
	r.Req.SetUser(toUser)
	r.Req.SetHost(toDomain)

	r.To.SetUriType("sip")
	r.To.SetUser(toUser)
	r.To.SetHost(toDomain)

	r.From.SetUriType("sip")
	r.From.SetUser(fromUser)
	r.From.SetHost(fromDomain)

	r.Cseq.SetID("1")
	r.Cseq.SetMethod(request)
	r.ContLen.SetValue("0")
	return r
}
