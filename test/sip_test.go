package test

import "testing"
import "github.com/KalbiProject/Kalbi/sip/message"

func TestSIPParser(t *testing.T) {
	byteMsg := []byte(msg)
	x := message.Parse(byteMsg)

	t.Log(string(x.Auth.QoP))
	t.Log(string(x.Auth.Nonce))
	t.Log(string(x.Auth.Realm))
	t.Log(string(x.Auth.Algorithm))
	t.Log(string(x.Auth.Username))
	t.Log(string(x.Auth.Nc))
	t.Log(string(x.Auth.CNonce))
	t.Log(string(x.Auth.Response))
	t.Log(string(x.Auth.URI))
	if string(x.Req.Method) != "INVITE" {
		t.Error("Method line not parsed")
	}
}
