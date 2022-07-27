package message

import (
	"fmt"
	"github.com/KalbiProject/kalbi/sip/status"
	"strconv"
)

// NewResponse creates new SIP Response
func NewResponse(tx Transaction, statuscode int, body []byte) *SipMsg {
	statusline := NewResponseLine(statuscode, status.StatusText(statuscode))
	response := new(SipMsg)
	response.Req = *statusline
	response.CopyHeaders(tx.GetOrigin())
	lp := tx.GetListeningPoint()
	response.Contact.SetHost(lp.GetHost())
	response.Contact.SetPort(strconv.Itoa(lp.GetPort()))
	response.Body = body
	response.ContLen.SetValue(fmt.Sprint(len(response.Body)))
	return response
}
