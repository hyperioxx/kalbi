package proxy

import (
	"fmt"
	//"strconv"
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/sip/status"
	"github.com/KalbiProject/Kalbi/sip/transaction"
	"strings"
)

type Proxy struct {
	stack            *kalbi.SipStack
	requestschannel  chan transaction.Transaction
	responseschannel chan transaction.Transaction
	RegisteredUsers  map[string]string
}

func (p *Proxy) HandleRequest(tx transaction.Transaction) {
	if string(tx.GetLastMessage().Req.Method) == method.CANCEL {
		msg := message.NewResponse(status.OK, string(tx.GetOrigin().To.User)+"@"+string(tx.GetOrigin().To.Host), string(tx.GetOrigin().From.User)+"@"+string(tx.GetOrigin().From.Host))
		msg.CopyHeaders(tx.GetOrigin())
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))
        
	}
	if string(tx.GetLastMessage().Req.Method) == method.INVITE {

		msg := message.NewResponse(status.Trying, string(tx.GetOrigin().Contact.Host)+"@"+string(tx.GetOrigin().Contact.Host), "@")
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

		user, exists := p.RegisteredUsers[string(tx.GetOrigin().To.User)]
		if exists == false {
			msg := message.NewResponse(status.NotFound, "@", "@")
			msg.CopyHeaders(tx.GetOrigin())
			tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

		} else {
			user := strings.Split(user, ":")
			msg2 := message.NewRequest(method.INVITE, string(tx.GetOrigin().To.User)+"@"+string(tx.GetOrigin().To.Host), string(tx.GetOrigin().From.User)+"@"+string(tx.GetOrigin().From.Host))
			msg2.CopyHeaders(tx.GetOrigin())
			msg2.CopySdp(tx.GetOrigin())
            msg2.To.Host = []byte(user[0])
			msg2.Via[0].Host = []byte("192.168.10.122")
			msg2.Via[0].SetBranch(transaction.GenerateBranchId())
			msg2.Contact.Host = []byte("192.168.10.122")
			msg2.Contact.Port = []byte("5060")
			msg2.Via[0].Port = []byte("5060")
			TxMng := p.stack.GetTransactionManager()
			ctx := TxMng.NewClientTransaction(msg)
			ctx.ServerTxID = string(tx.GetBranchID())
			fmt.Println("Branch: " + string(tx.GetLastMessage().Via[0].Branch))
			fmt.Println(msg2.Sdp.Export())
			fmt.Printf("SDP SIZE: %d", len(msg2.Sdp.Export()))
            //msg2.ContLen.SetValue(strconv.Itoa(msg.Sdp.Size()))
			ctx.Send(msg2, user[0], user[1])

		}

	} else if string(tx.GetOrigin().Req.Method) == method.REGISTER {
		p.RegisteredUsers[string(tx.GetOrigin().Contact.User)] = string(tx.GetOrigin().Contact.Host) + ":" + string(tx.GetOrigin().Contact.Port)
		msg := message.NewResponse(status.OK, "@", "@")
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	} else if string(tx.GetOrigin().Req.Method) == method.BYE {
		msg := message.NewResponse(status.OK, "@", "@")
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	} else {
		msg := message.NewResponse(status.OK, "@", "@")
		msg.CopyHeaders(tx.GetOrigin())
		msg.ContLen.SetValue("0")
		tx.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	}

}


func (p *Proxy) HandleResponse(response transaction.Transaction) {

	 
	 if response.GetLastMessage().GetStatusCode() == 100 {
		 return

	} else if response.GetLastMessage().GetStatusCode() == 180 {
		 return
	
	 } else {
		  
		  TxMng := p.stack.GetTransactionManager()
		  tx, exists := TxMng.FindServerTransactionByID(response.GetServerTransactionID())
		  if exists == false {
			  fmt.Println("Server Tranaction has been terminated as 200 OK has been sent, must send statlessly ")
			  msg := message.NewResponse(status.OK, "@", "@")
			  msg.CopyHeaders(response.GetOrigin())
			  msg.CopySdp(response.GetOrigin())
			  response.Send(msg, string(response.GetLastMessage().Contact.Host), string(response.GetLastMessage().Contact.Port))
		  }else {
			fmt.Println("Server Tranaction Found")
			msg := message.NewResponse(status.OK, "@", "@")

			fmt.Println(string(response.GetLastMessage().Src))
			msg.CopyHeaders(response.GetLastMessage())
			msg.Via[0].SetBranch(tx.GetBranchID())
			msg.CopySdp(response.GetLastMessage())
			//msg.ContLen.SetValue(strconv.Itoa(msg.Sdp.Size()))
			fmt.Println(string(msg.Export()))
			response.Send(msg, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

		  }
	
	 }

    
}

func (p *Proxy) AddToRegister(key string, contact string) {
	p.RegisteredUsers[key] = contact
}

func (p *Proxy) ServeRequests() {

	for {
		tx := <-p.requestschannel
		p.HandleRequest(tx)
	}

}

func (p *Proxy) ServeResponses() {
	for {
		tx := <-p.responseschannel
		p.HandleResponse(tx)
	}
}

func (p *Proxy) Start() {
	p.RegisteredUsers = make(map[string]string)
	p.stack = kalbi.NewSipStack("Basic")
	p.stack.CreateListenPoint("udp", "192.168.10.122", 5060)
	//p.stack.CreateListenPoint("udp", "127.0.0.1", 5060)
	p.requestschannel = p.stack.CreateRequestsChannel()
	p.responseschannel = p.stack.CreateResponseChannel()
	go p.stack.Start()
	go p.ServeRequests()
	
	
	p.ServeResponses()

}
