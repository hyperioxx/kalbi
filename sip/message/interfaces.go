package message


//SipListener interface for sip listener i.e. Your Application
type SipListener interface {
	HandleRequests(SipEventObject)
	HandleResponses(SipEventObject)
}


//Transaction interface for SIP transactions
type Transaction interface {
	GetBranchID() string
	GetOrigin() *SipMsg
	SetListeningPoint(ListeningPoint)
	Send(*SipMsg, string, string)
	Receive(*SipMsg)
	GetLastMessage() *SipMsg
	GetServerTransactionID() string
	SetLastMessage(*SipMsg)
	GetListeningPoint() ListeningPoint
}

//SipEventObject interface for SIP events
type SipEventObject interface {
	GetSipMessage() *SipMsg
	SetSipMessage(*SipMsg)
	GetTransaction() Transaction
	SetTransaction(Transaction)
	SetListeningPoint(ListeningPoint)
	GetListeningPoint() ListeningPoint
}

//ListeningPoint interface for listening point
type ListeningPoint interface {
	Read() SipEventObject
	Build(string, int)
	Start()
	GetHost() string
	GetPort() int
	SetTransportChannel(chan SipEventObject)
	Send(string, string, string) error
}
