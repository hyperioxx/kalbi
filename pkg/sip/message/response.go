package message


import ("strings")



func NewResponse(response string, to string , from string) *SipMsg {
	to_ := strings.Split(to, "@")
	from_ := strings.Split(from, "@")
	_response := strings.Split(response, " ")
	r := new(SipMsg)
	r.Req = *new(SipReq)
	r.To = *new(SipTo)
	r.From = *new(SipFrom)

	r.Req.SetStatusCode(_response[0])
	r.Req.SetStatusDesc(_response[1])
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







