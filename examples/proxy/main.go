package main

import (
	"fmt"
	"Kalbi/pkg/sip/stack"
	"Kalbi/pkg/sip/events"
)


type Proxy struct {
	stack *stack.SipStack
	eventschannel chan events.Event

}


func (p *Proxy) HandleRequest(request events.Event){
	fmt.Println("I am handling a request")
	
}


func (p *Proxy) HandleResponse(response events.Event){
	fmt.Println("I am handling a response")
}




func (p *Proxy) Start() {
        p.stack = stack.NewSipStack("Basic")
		p.stack.CreateListenPoint("udp", "0.0.0.0", 5060)
		p.eventschannel = p.stack.CreateEventsChannel()
		go p.stack.Start()

		for{
		   event := <-p.eventschannel
		   if event.Type == "Server"{
			   p.HandleRequest(event)
		   }else if event.Type == "Client"{
			   p.HandleResponse(event)
		   }
		}
}



func main(){
   proxy := new(Proxy)
   proxy.Start()
}