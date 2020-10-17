package event

import (
	"github.com/KalbiProject/Kalbi/interfaces"
	"github.com/KalbiProject/Kalbi/sip/message"
)

//SipEvent object that gets passed to the SipListener
type SipEvent struct {
	sipmsg *message.SipMsg
	tx     interfaces.Transaction
}

//GetSipMessage returns message that created this event
func (se *SipEvent) GetSipMessage() *message.SipMsg {
	return se.sipmsg
}

//SetSipMessage sets message that created this event
func (se *SipEvent) SetSipMessage(msg *message.SipMsg) {
	se.sipmsg = msg
}

//GetTransaction returns transaction related to the SIP message that created this event
func (se *SipEvent) GetTransaction() interfaces.Transaction {
	return se.tx
}

//SetTransaction sets transaction related to the SIP message that created this event
func (se *SipEvent) SetTransaction(tx interfaces.Transaction) {
	se.tx = tx
}
