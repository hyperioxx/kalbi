package main

import (
	"fmt"
	"Kalbi/pkg/sip/stack"
	"Kalbi/pkg/sip/status"
	"Kalbi/pkg/sip/method"
	"Kalbi/pkg/sip/transaction"
    "Kalbi/pkg/sip/message"
)


type Proxy struct {
	stack *stack.SipStack
	requestschannel chan transaction.Transaction
	responseschannel chan transaction.Transaction
	
}


func (p *Proxy) HandleRequest(tx transaction.Transaction){
	if string(tx.GetOrigin().Req.Method) == method.INVITE{
		msg := message.NewResponse(status.TRYING_100, "@", "@")
		msg.CopyMessage(tx.GetOrigin())
		msg.ContLen.SetValue("0")
        tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

        msg2 := message.NewResponse(status.OK_200, "@", "@")
		msg2.CopyMessage(tx.GetOrigin())
		msg2.ContLen.SetValue("0")

		tx.Send(msg2, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	}else if string(tx.GetOrigin().Req.Method) == method.REGISTER{
		msg := message.NewResponse(status.OK_200, "@", "@")
		msg.CopyMessage(tx.GetOrigin())
		
		msg.ContLen.SetValue("0")

		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	}else if string(tx.GetOrigin().Req.Method) == method.BYE{
		msg := message.NewResponse(status.OK_200, "@", "@")
		msg.CopyMessage(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))
	}

}


func (p *Proxy) HandleResponse(response transaction.Transaction){
	fmt.Println(string(response.GetOrigin().Req.Src))
	
}

func(p *Proxy) ServeRequests(){

	for{
		tx := <- p.requestschannel
		p.HandleRequest(tx)
	}


}

func (p *Proxy) ServeResponses(){
	for {
	    tx := <-p.responseschannel
        p.HandleResponse(tx)
    }
}




func (p *Proxy) Start() {
        p.stack = stack.NewSipStack("Basic")
		p.stack.CreateListenPoint("udp", "0.0.0.0", 5060)
		p.requestschannel = p.stack.CreateRequestsChannel()
		p.responseschannel = p.stack.CreateResponseChannel()
		go p.stack.Start()
		go p.ServeRequests()
		p.ServeResponses()
	
}



func main(){
   proxy := new(Proxy)
   proxy.Start()
}