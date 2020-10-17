package test

import "testing"
import "github.com/KalbiProject/Kalbi/sdp"

func TestSDPParser(t *testing.T) {
	byteMsg := []byte(MSG)
	x := sdp.Parse(byteMsg)
	t.Log(string(x.Time.String()))

}
