package application

import "fmt"
import "Kalbi/sip/message"

//B2BUA is a sip proxy
type B2BUA struct {

}

func (b *B2BUA) HandleRequest(req *message.Request){
	fmt.Println(req.To.Hostname)
}