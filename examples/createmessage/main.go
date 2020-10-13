package main

import (
	"fmt"
	"github.com/KalbiProject/Kalbi"
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/sip/method"
	"github.com/KalbiProject/Kalbi/sip/status"
)

func main() {
	x := message.NewRequest(method.INVITE, "12345@127.0.0.1", "4321@127.0.0.1")
	fmt.Println(x.Export())

	y := message.NewResponse(status.OK, "12345@127.0.0.1", "4321@127.0.0.1")
	fmt.Println(y.Export())

}
