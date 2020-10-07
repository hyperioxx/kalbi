package main

import (
	"fmt"
	"Kalbi/pkg/sip/stack"
	"Kalbi/pkg/sip/status"
	"Kalbi/pkg/sip/method"
	"Kalbi/pkg/sip/transaction"
    "Kalbi/pkg/sip/message"
)


type SipClient struct {
	stack *stack.SipStack
	requestschannel chan transaction.Transaction
	responseschannel chan transaction.Transaction
}


func (p *SipClient) HandleRequest(tx transaction.Transaction){
	

}


func (p *SipClient) HandleResponse(response transaction.Transaction){
	fmt.Println(string(response.GetOrigin().Req.Src))
	
}

func(p *SipClient) ServeRequests(){

	for{
		tx := <- p.requestschannel
		p.HandleRequest(tx)
	}


}

func (p *SipClient) ServeResponses(){
	for {
	    tx := <-p.responseschannel
        p.HandleResponse(tx)
    }
}


func (p *SipClient) SendInvite(){
	msg := message.NewRequest(method.INVITE, "123@127.0.0.1", "321@127.0.0.1")
	txmng := p.stack.GetTransactionManager()
	tx := txmng.NewClientTransaction(msg)
	tx.Send(msg, "127.0.0.1", "5060")
}


func (p *SipClient) Start() {
        p.stack = stack.NewSipStack("Basic")
		p.stack.CreateListenPoint("udp", "0.0.0.0", 5060)
		p.requestschannel = p.stack.CreateRequestsChannel()
		p.responseschannel = p.stack.CreateResponseChannel()
		go p.stack.Start()
		go p.ServeRequests()
		go p.ServeResponses()
	
}



func main(){
   proxy := new(SipClient)
   proxy.Start()
   

}