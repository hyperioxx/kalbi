package sip

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/KalbiProject/kalbi/sdp"
	"github.com/KalbiProject/kalbi/sip/message"
)

var (
	testInBytes = []byte{116, 101, 115, 116}
	testOutput  = message.SipMsg{
		Req: message.SipReq{
			Method:     []uint8(nil),
			UriType:    "",
			StatusCode: []uint8(nil),
			StatusDesc: []uint8(nil),
			User:       []uint8(nil),
			Host:       []uint8(nil),
			Port:       []uint8(nil),
			UserType:   []uint8(nil),
			Src:        []uint8{0x74, 0x65, 0x73, 0x74}},
		From: message.SipFrom{
			UriType:  "",
			Name:     []uint8(nil),
			User:     []uint8(nil),
			Host:     []uint8(nil),
			Port:     []uint8(nil),
			Tag:      []uint8(nil),
			UserType: []uint8(nil),
			Src:      []uint8(nil)},
		To: message.SipTo{UriType: "",
			Name:     []uint8(nil),
			User:     []uint8(nil),
			Host:     []uint8(nil),
			Port:     []uint8(nil),
			Tag:      []uint8(nil),
			UserType: []uint8(nil),
			Src:      []uint8(nil)},
		Contact: message.SipContact{
			UriType: "",
			Name:    []uint8(nil),
			User:    []uint8(nil),
			Host:    []uint8(nil),
			Port:    []uint8(nil),
			Tran:    []uint8(nil),
			Qval:    []uint8(nil),
			Expires: []uint8(nil),
			Src:     []uint8(nil)},
		Via: []message.SipVia{},
		Cseq: message.SipCseq{
			ID:     []uint8(nil),
			Method: []uint8(nil),
			Src:    []uint8(nil)},
		Auth: message.SipAuth{
			Username:  []uint8(nil),
			Realm:     []uint8(nil),
			Nonce:     []uint8(nil),
			CNonce:    []uint8(nil),
			QoP:       []uint8(nil),
			Algorithm: []uint8(nil),
			Nc:        []uint8(nil),
			URI:       []uint8(nil),
			Response:  []uint8(nil),
			Src:       []uint8(nil)},
		Ua: message.SipVal{
			Value: []uint8(nil),
			Src:   []uint8(nil)},
		Exp: message.SipVal{
			Value: []uint8(nil),
			Src:   []uint8(nil)},
		MaxFwd: message.SipVal{
			Value: []uint8(nil),
			Src:   []uint8(nil)},
		CallID: message.SipVal{
			Value: []uint8(nil),
			Src:   []uint8(nil)},
		ContType: message.SipVal{
			Value: []uint8(nil),
			Src:   []uint8(nil)},
		ContLen: message.SipVal{
			Value: []uint8(nil),
			Src:   []uint8(nil)},
		Src:  []uint8{0x74, 0x65, 0x73, 0x74},
		Body: []uint8(nil),
		Sdp: sdp.SdpMsg{
			Origin: sdp.SdpOrigin{
				Username:       []uint8(nil),
				SessionID:      []uint8(nil),
				SessionVersion: []uint8(nil),
				NetType:        []uint8(nil),
				AddrType:       []uint8(nil),
				UniAddr:        []uint8(nil),
				Src:            []uint8(nil)},
		}}
)

func TestParse(t *testing.T) {
	type args struct {
		msg []byte
	}
	tests := []struct {
		name string
		args args
		want message.SipMsg
	}{
		{"test Parse", args{msg: testInBytes}, testOutput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got.Src)
				fmt.Printf("%#v", got)
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
