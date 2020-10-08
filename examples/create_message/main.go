package main

import "fmt"
import "Kalbi/sip/status"
import "Kalbi/sip/method"
import "Kalbi/sip/message"



func main(){
	x := message.NewRequest(method.INVITE, "12345@127.0.0.1", "4321@127.0.0.1")
	fmt.Println(x.Export())

	y := message.NewResponse(status.OK_200, "12345@127.0.0.1", "4321@127.0.0.1")
	fmt.Println(y.Export())

}