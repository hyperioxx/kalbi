package main

import (
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/interfaces"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/sip/status"
)

type Client struct {
	stack *kalbi.SipStack
}

func (c *Client) HandleRequests(event interfaces.SipEventObject) {

	tx := event.GetTransaction()
	switch string(tx.GetLastMessage().Req.Method) {
	case method.CANCEL:
		//handle CANCEL request
	case method.INVITE:
		//handle INVITE request
	case method.REGISTER:
		//handle REGISTER request
	case method.BYE:
		//handle BYE request
	case method.ACK:

	default:
		msg := message.NewResponse(status.OK, "@", "@")
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))
	}

}

func (c *Client) HandleResponses(event interfaces.SipEventObject) {

	response := event.GetTransaction()

	switch response.GetLastMessage().GetStatusCode() {

	case 100:
		//Handle 100 Trying
	case 180:
		//Handle 180 Ringing
	case 200:
		//Handle 200 OK
	default:
		//Handle Default
	}

}

func (c *Client) SendInvite() {
	request := message.NewRequest(method.INVITE, "4321@127.0.0.1", "1234@127.0.0.1") //Args:  METHOD  TO HEADER  FROM HEADER
}

func (c *Client) Start(host string, port int) {
	c.stack = kalbi.NewSipStack("Basic Client Example")
	c.stack.CreateListenPoint("udp", host, port)
	c.stack.SetSipListener(c)
	go c.stack.Start()
}

func main() {
	client := new(Client)
	client.Start("127.0.0.1", 5060)
	client.SendInvite()
	select {} //blocking action
}
