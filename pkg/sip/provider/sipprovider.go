package sipprovider

import "Kalbi/pkg/transport"

func NewSipProvider() *SipProvider {
	sp := new(SipProvider)
	sp.ListeningPoint 
}


type SipProvider struct {
	ListeningPoint  
	
}

//GetListeningPoint Returns the ListeningPoint of this SipProvider
func (sp *SipProvider) GetListeningPoint(){

}

//GetNewCallID Returns a unique CallIdHeader for identifying dialogues between two SIP applications.
func (sp *SipProvider) GetNewCallID(){

}


