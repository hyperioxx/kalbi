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
	"fmt"
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/interfaces"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/sip/status"
)

const title = `##    ##    ###    ##       ########  ####     ######  ##       #### ######## ##    ## ######## 
##   ##    ## ##   ##       ##     ##  ##     ##    ## ##        ##  ##       ###   ##    ##    
##  ##    ##   ##  ##       ##     ##  ##     ##       ##        ##  ##       ####  ##    ##    
#####    ##     ## ##       ########   ##     ##       ##        ##  ######   ## ## ##    ##    
##  ##   ######### ##       ##     ##  ##     ##       ##        ##  ##       ##  ####    ##    
##   ##  ##     ## ##       ##     ##  ##     ##    ## ##        ##  ##       ##   ###    ##    
##    ## ##     ## ######## ########  ####     ######  ######## #### ######## ##    ##    ##    `

type ClientProperties struct {
	IP          string
	Username    string
	Domain      string
	Password    string
	Registrar   string
}

type Client struct {
	stack       *kalbi.SipStack
	properties  *ClientProperties
}


func (c *Client) HandleRequests(event interfaces.SipEventObject) {

	tx := event.GetTransaction()
	switch string(tx.GetLastMessage().Req.Method) {
	case method.CANCEL:
		//handle CANCEL request
	case method.INVITE:
		responseLine := message.NewResponseLine(status.OK, "It's Cool")
		msg := message.NewResponse(responseLine, nil, nil, nil, nil, nil)
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	case method.REGISTER:
		//handle REGISTER request
	case method.BYE:
		//handle BYE request
	case method.ACK:

	default:
		responseLine := message.NewResponseLine(status.OK, "It's Cool")
		msg := message.NewResponse(responseLine, nil, nil, nil, nil, nil)
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))
	}

}

func (c *Client) HandleResponses(event interfaces.SipEventObject) {

	response := event.GetTransaction()
    fmt.Println(string(event.GetSipMessage().Src))
	switch response.GetLastMessage().GetStatusCode() {

	case 100:
		//Handle 100 Trying
	case 180:
		//Handle 180 Ringing
	case 200:
		//Handle 200 OK
	case 401:
        
	default:
		//Handle Default
	}

}
func (c *Client) SendRegister() {

	requestLine := message.NewRequestLine(method.REGISTER, "sip", c.properties.Username, c.properties.Domain, "5060") //Create requestline e.g.  REGISTER sip:1234@127.0.0.1:5060 SIP/2.0
	requestVia := message.NewViaHeader("udp", c.properties.IP, "5060") //Creates Via e.g. Via: SIP/2.0/UDP 127.0.0.1:5060
	requestVia.SetBranch(message.GenerateBranchId())
	requestFrom := message.NewFromHeader(c.properties.Username, "sip", c.properties.Domain, "5060") //Creates From e.g. From: <sip:1234@127.0.0.1>
	requestTo := message.NewToHeader(c.properties.Username, "sip", c.properties.Domain, "5060") //Creates To e.g. To: <sip:5678@127.0.0.1>
	requestContact := message.NewContactHeader("sip", c.properties.Username, c.properties.IP)
	requestCallID := message.NewCallID("123456789")//Creates CallID e.g. Call-ID: 123456789
	requestCseq := message.NewCSeq("1", method.REGISTER) //Creates CSeq e.g. CSeq: 1 INVITE
	requestMaxFor := message.NewMaxForwards("70") //Creates Max-Forwards e.g. Max-Forwards: 70
	requestContentLen := message.NewContentLength("0")
	request := message.NewRequest(requestLine, requestVia , requestTo, requestFrom, requestContact, requestCallID, requestCseq, requestMaxFor, requestContentLen)

	txmng := c.stack.GetTransactionManager()
	txmng.NewClientTransaction(request)
	c.stack.ListeningPoints[0].Send(c.properties.Registrar, "5060", request.String())
}

func (c *Client) Start(host string, port int) {
	c.stack = kalbi.NewSipStack("Basic Client Example")
	c.stack.CreateListenPoint("udp", host, port)
	c.stack.SetSipListener(c)
	go c.stack.Start()
}



func configure() *ClientProperties {
	props := new(ClientProperties)
	fmt.Println(title + "\n\n\n\n")

	//Username
	fmt.Print("Username:")
	_, err := fmt.Scan(&props.Username)
	if err != nil {
		fmt.Println(err)
	}

	//Domain
	fmt.Print("Domain:")
	_, err = fmt.Scan(&props.Domain)
	if err != nil {
		fmt.Println(err)
	}

	//Password will be visable because I cba to set stty echo off using syscalls 
	fmt.Print("Password:")
	_, err = fmt.Scan(&props.Password)
	if err != nil {
		fmt.Println(err)
	}

	//IP 
	fmt.Print("Machine IP:")
	_, err = fmt.Scan(&props.IP)
	if err != nil {
		fmt.Println(err)
	}

	//Registrar
	fmt.Print("Registrar:")
	_, err = fmt.Scan(&props.Registrar)
	if err != nil {
		fmt.Println(err)
	}

	return props
}


func basicCliInterface() {

	alive := true
	var command string
	for alive == true{
		
		fmt.Print("kalbiclient> ")

		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println(err)
		}
		if command == "exit"{
             alive = false
		}else {
			fmt.Println("Unknown command")
		}
		
	}

}



func main() {
	props := configure()
	client := new(Client)
	client.properties = props
	client.Start(props.IP, 5060)
	client.SendRegister()
	basicCliInterface() 
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