package test

import (
	"fmt"
	"testing"
	"github.com/KalbiProject/Kalbi/sip/message"
)

func TestSIPParser(t *testing.T) {
	byteMsg := []byte(MSG)
	x := message.Parse(byteMsg)

    fmt.Println(string(x.Body))
	if string(x.Req.Method) != "INVITE" {
		t.Error("Method line not parsed")
	}
}
