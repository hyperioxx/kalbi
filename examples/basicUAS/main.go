package main

import (
	kalbi "github.com/KalbiProject/kalbi"
	"github.com/KalbiProject/kalbi/sip/message"
	"github.com/KalbiProject/kalbi/sip/status"
)

func main() {
	stack := kalbi.NewSipStack("My New Sip Stack")

	stack.REGISTER(func(event message.SipEventObject) {
		tx := event.GetTransaction()
		response := message.NewResponse(tx, status.OK, nil)
		tx.Send(response, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	})

	stack.INVITE(func(event message.SipEventObject) {
		tx := event.GetTransaction()
		response := message.NewResponse(tx, status.OK, nil)
		tx.Send(response, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	})

	stack.BYE(func(event message.SipEventObject) {
		tx := event.GetTransaction()
		response := message.NewResponse(tx, status.OK, nil)
		tx.Send(response, string(tx.GetOrigin().Contact.Host), string(tx.GetOrigin().Contact.Port))

	})

	stack.ACK(func(event message.SipEventObject) {
		_ = event.GetTransaction()

	})

	stack.CreateListenPoint("udp", "127.0.0.1", 5060)
	stack.CreateListenPoint("udp", "127.0.0.1", 5061)
	stack.Start()
}
