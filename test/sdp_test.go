package test

import (
	"testing"

	"github.com/KalbiProject/kalbi/sdp"
)

func TestSDPParser(t *testing.T) {
	byteMsg := []byte(msg)
	x := sdp.Parse(byteMsg)
	t.Log(string(x.Time.String()))

}
