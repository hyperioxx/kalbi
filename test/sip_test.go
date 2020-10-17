package test

import (
	"fmt"
	"github.com/KalbiProject/Kalbi/sip/message"
	"testing"
)

func TestSIPParser(t *testing.T) {
	byteMsg := []byte(MSG)
	x := message.Parse(byteMsg)

	fmt.Println(string(x.Body))
	if string(x.Req.Method) != "INVITE" {
		t.Error("Method line not parsed")
	}
}
