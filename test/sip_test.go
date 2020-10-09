package test


import "testing"
import "github.com/KalbiProject/Kalbi/sip/message"



func TestSIPParser(t *testing.T){
	byteMsg := []byte(msg) 
	x := message.Parse(byteMsg)

	if string(x.Req.Method) != "INVITE" {
         t.Error("Method line not parsed" )
	}
}