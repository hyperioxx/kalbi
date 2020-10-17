package event

import ("github.com/KalbiProject/Kalbi/sip/message"
"github.com/KalbiProject/Kalbi/interfaces")


type ListeningPoint interface {
	Read() *message.SipMsg
	Build(string, int)
	Start()
	SetTransportChannel(chan *message.SipMsg)
	Send(string, string, string) error
}

type Transaction interface {
	GetBranchID() string
	GetOrigin() *message.SipMsg
	SetListeningPoint(ListeningPoint)
	Send(*message.SipMsg, string, string)
	Receive(*message.SipMsg)
	GetLastMessage() *message.SipMsg
	GetServerTransactionID() string
   SetLastMessage(*message.SipMsg)
   GetListeningPoint() ListeningPoint
}

type SipEventObject interface{
	GetSipMessage() *message.SipMsg
	SetSipMessage(*message.SipMsg)
	GetTransaction() interfaces.Transaction
	SetTransaction( interfaces.Transaction)
}


type SipEvent struct {
	sipmsg    *message.SipMsg
    tx        interfaces.Transaction
    
}


func (se * SipEvent) GetSipMessage() *message.SipMsg {
     return se.sipmsg
}

func (se *SipEvent) SetSipMessage(msg *message.SipMsg) {
	se.sipmsg = msg
}

func (se *SipEvent) GetTransaction() interfaces.Transaction {
	return se.tx
}

func (se *SipEvent) SetTransaction(tx interfaces.Transaction) {
	se.tx = tx
}