package main

import (
	"fmt"
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/sip/status"
	"github.com/KalbiProject/Kalbi/sip/transaction"
)

type SipClient struct {
	stack            *kalbi.SipStack
	requestschannel  chan transaction.Transaction
	responseschannel chan transaction.Transaction
}

func (p *SipClient) HandleRequest(tx transaction.Transaction) {

}

func (p *SipClient) HandleResponse(response transaction.Transaction) {
	fmt.Println(string(response.GetOrigin().Req.Src))

}

func (p *SipClient) ServeRequests() {

	for {
		tx := <-p.requestschannel
		p.HandleRequest(tx)
	}

}

func (p *SipClient) ServeResponses() {

	msg := message.NewRequest(method.INVITE, "123@127.0.0.1", "321@127.0.0.1")
	fmt.Println(msg)
	txmng := p.stack.GetTransactionManager()
	fmt.Println(txmng)
	tx := txmng.NewClientTransaction(msg)
	fmt.Println(tx)
	tx.Send(msg, "193.168.10.138", "5060")

	for {
		tx := <-p.responseschannel
		p.HandleResponse(tx)
	}
}

func (p *SipClient) SendInvite() {
	msg := message.NewRequest(method.INVITE, "123@127.0.0.1", "321@127.0.0.1")
	txmng := p.stack.GetTransactionManager()
	tx := txmng.NewClientTransaction(msg)
	tx.Send(msg, "127.0.0.1", "5060")
}

func (p *SipClient) Start() {
	p.stack = kalbi.NewSipStack("Basic")
	p.stack.CreateListenPoint("udp", "0.0.0.0", 5060)
	p.requestschannel = p.stack.CreateRequestsChannel()
	p.responseschannel = p.stack.CreateResponseChannel()
	go p.stack.Start()
	go p.ServeRequests()
	p.ServeResponses()

}

func main() {
	proxy := new(SipClient)
	proxy.Start()

}
