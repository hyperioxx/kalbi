package siplistener 

type SipListener interface{
	ProcessRequest()
	ProcessResponse()
}