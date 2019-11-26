package application

import "Kalbi/sip/message"

//Manager manages sip applications
type Manager struct {
	input chan *message.Request
	
}

//SetChannel sets channel
func (am *Manager) SetChannel(input chan *message.Request){
    am.input = input
}

//Start starts application manager
func (am *Manager) Start(){
	registrar := new(Registrar)
	b2bua := new(B2BUA)
	for {
		request := <- am.input
		
		switch request.Method{
		case "REGISTER":
			registrar.HandleRequest(request)
		default:
			b2bua.HandleRequest(request)
		}

	}


}


//NewAppManager creates new App Manager
func NewAppManager(input chan *message.Request) *Manager{
	appManager := new(Manager)
	appManager.SetChannel(input)
	return appManager
}