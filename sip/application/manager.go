package application

import "fmt"
import "Kalbi/sip/message"

//ApplicationManager manages sip applications
type ApplicationManager struct {
     input chan *message.Request
}

//SetChannels sets channel
func (am *ApplicationManager) SetChannel(input chan *message.Request){
    am.input = input
}

//Start starts application manager
func (am *ApplicationManager) Start(){
	for {
		request := <- am.input
		
		switch request.Method{
		case "REGISTER":
			fmt.Println("Handle REGISTER")
		default:
			fmt.Println("HANDLE")

		}

	}


}


//NewAppManager creates new App Manager
func NewAppManager(input chan *message.Request) *ApplicationManager{
	appManager := new(ApplicationManager)
	appManager.SetChannel(input)
	return appManager
}