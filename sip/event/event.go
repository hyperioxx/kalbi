package event

import (
	"kalbi/sip/message"
)

// SipEvent object that gets passed to the SipListener
type SipEvent struct {
	sipmsg *message.SipMsg
	tx     message.Transaction
	lp     message.ListeningPoint
}

// GetSipMessage returns message that created this event
func (se *SipEvent) GetSipMessage() *message.SipMsg {
	return se.sipmsg
}

// SetSipMessage sets message that created this event
func (se *SipEvent) SetSipMessage(msg *message.SipMsg) {
	se.sipmsg = msg
}

// GetTransaction returns transaction related to the SIP message that created this event
func (se *SipEvent) GetTransaction() message.Transaction {
	return se.tx
}

// SetTransaction sets transaction related to the SIP message that created this event
func (se *SipEvent) SetTransaction(tx message.Transaction) {
	se.tx = tx
}

// SetListeningPoint gives ability to set message.ListeningPoint
func (se *SipEvent) SetListeningPoint(lp message.ListeningPoint) {
	se.lp = lp
}

// GetListeningPoint returns message.ListeningPoint
func (se *SipEvent) GetListeningPoint() message.ListeningPoint {
	return se.lp
}
