package main

import (
	"fmt"

	kalbi "github.com/KalbiProject/kalbi"
	"github.com/KalbiProject/kalbi/interfaces"
)

func main() {
	stack := kalbi.NewSipStack("My New Sip Stack")

	stack.REGISTER(func(event interfaces.SipEventObject) {

		tx := event.GetTransaction()
		fmt.Println(tx.GetLastMessage())
		//tx.Send()

	})
	stack.INVITE(func(event interfaces.SipEventObject) {

		tx := event.GetTransaction()
		fmt.Println(tx.GetLastMessage())
		//tx.Send()

	})

	stack.CreateListenPoint("udp", "127.0.0.1", 5060)
	stack.CreateListenPoint("udp", "127.0.0.1", 5061)
	stack.Start()
}
