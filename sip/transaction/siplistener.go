package transaction



type SipListener interface {
	HandleRequests(Transaction)
	HandleResponses(Transaction)
}