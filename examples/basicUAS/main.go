package main

import (
	//"fmt"
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/interfaces"
)


func main() {
	stack := kalbi.NewSipStack("My New Sip Stack")

	stack.REGISTER(func(event interfaces.SipEventObject) {
	
		tx := event.GetTransaction()
		tx.Send()

	})

	stack.CreateListenPoint("udp", "0.0.0.0", 5060)
	stack.CreateListenPoint("udp", "0.0.0.0", 5061)
	stack.Start()
}