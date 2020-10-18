![Kalbi Logo](https://raw.githubusercontent.com/hyperioxx/Kalbi/master/assets/images/logo_transparent_background.png "Kalbi Logo")




<h1 align="center">
  Golang Session Initiated Protocol Framework/SDK
</h1>

<h3 align="center">
  <a href="https://pkg.go.dev/github.com/KalbiProject/Kalbi">Documentation</a> • 
  <a href="https://twitter.com/KalbiProject">Twitter</a> • 
  <a href="https://www.reddit.com/r/Kalbi/">Community</a> •
  <a href="https://discord.gg/6NCKgrz">Discord</a>
</h3>

<p>&nbsp;</p>



 [![Go Report Card](https://goreportcard.com/badge/github.com/KalbiProject/Kalbi)](https://goreportcard.com/report/github.com/KalbiProject/Kalbi) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/KalbiProject/Kalbi/LICENCE)  [![GitHub contributors](https://img.shields.io/github/contributors/KalbiProject/Kalbi)](https://github.com/KalbiProject/Kalbi/graphs/contributors/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Hyperioxx/Kalbi) 
<br />



## Description

Golang SIP/VoIP SDK/Framework used to build large platforms for VoIP and realtime communications



## Examples

[KalbiProxy](https://github.com/KalbiProject/KalbiProxy) - KalbiProxy is a SIP Proxy/Registrar Server built using the Kalbi Framework/SDK


## Basic Example

Still in development

```golang

package main


import (
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/interfaces"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/sip/status"
)

type Client struct {
	stack            *kalbi.SipStack
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


func (c *Client) SendInvite(){
	  request := message.NewRequest(method.INVITE, "4321@127.0.0.1", "1234@127.0.0.1")//Args:  METHOD  TO HEADER  FROM HEADER
}


func (c *Client) Start(host string, port int) {
	c.stack = kalbi.NewSipStack("Basic Client Example")
	c.stack.CreateListenPoint("udp", host, port)
	c.stack.SetSipListener(p)
	go c.stack.Start()
}


func main(){
	client := new(Client)
	client.Start()
	client.SendInvite()
	select{} //blocking action 
}

```


### Core Team

<table>
   <tr>
      <td>
         <a href="https://github.com/hyperioxx"><img width="160px" src="https://avatars0.githubusercontent.com/u/17745250?s=400&u=561eac60ef16400408dc29f10ef36de8dbf011f9&v=4"><br>
         Aaron Parfitt</a><br>
        </td>
      <td>
         <a href="https://github.com/DeWarner"><img width="160px" src="https://avatars1.githubusercontent.com/u/20417324?s=460&u=42c60bbaa4a38e60394a1b9aeeb42dfd3969e708&v=4"><br>
         Declan Warner</a><br>  
      </td>
   </tr>
</table>