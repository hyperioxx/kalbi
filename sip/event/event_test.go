package event

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/KalbiProject/kalbi/interfaces"
	"github.com/KalbiProject/kalbi/sdp"
	"github.com/KalbiProject/kalbi/sip/message"
)

var (
	testSipVia  = make([]message.SipVia, 0)
	testSipMsg1 = message.SipMsg{
		Req: message.SipReq{
			Method:     []byte("INVITE"),
			UriType:    "sip",
			StatusCode: []byte(""),
			StatusDesc: []byte(""),
			User:       []byte("0"),
			Host:       []byte("test"),
			Port:       []byte(""),
			UserType:   []byte(""),
			Src:        []byte("INVITE sip:0@test SIP/2.0"),
		}, // SipReq
		From: message.SipFrom{
			UriType:  "sip",
			Name:     []byte("test"),
			User:     []byte("0"),
			Host:     []byte("test"),
			Port:     []byte("test"),
			Tag:      []byte("e788ae25565649ee9ea69e5fc09ec4"),
			UserType: []byte(""),
			Src:      []byte(`\"test\" <sip:test@test>;tag=5627f2c27f324ac28fe114c0ef620404"`),
		}, // SipFrom
		To: message.SipTo{
			UriType:  "sip",
			Name:     []byte(""),
			User:     []byte("0"),
			Host:     []byte("test"),
			Port:     []byte(""),
			Tag:      []byte(""),
			UserType: []byte(""),
			Src:      []byte("<sip:0@test>"),
		}, // SipTo
		Contact: message.SipContact{
			UriType: "sip",
			Name:    []byte("test"),
			User:    []byte("test"),
			Host:    []byte("127.0.0.1"),
			Port:    []byte("51231"),
			Tran:    []byte(""),
			Qval:    []byte(""),
			Expires: []byte(""),
			Src:     []byte(`test\" <sip:test@127.0.0.1:51231;ob>`),
		}, // SipContact
		Via: testSipVia, // []SipVia
		Cseq: message.SipCseq{
			ID:     []byte(""),
			Method: []byte(""),
		}, // SipCseq
		Auth:     message.SipAuth{}, // SipAuth
		Ua:       message.SipVal{},  // SipVal
		Exp:      message.SipVal{},  // SipVal
		MaxFwd:   message.SipVal{},  // SipVal
		CallID:   message.SipVal{},  // SipVal
		ContType: message.SipVal{},  // SipVal
		ContLen:  message.SipVal{},  // SipVal
		Src:      []byte(""),        // []byte
		Body:     []byte(""),        // []byte
		Sdp:      sdp.SdpMsg{},      // sdp.SdpMsg
	}

	testSipMsg2 = message.SipMsg{
		Req: message.SipReq{
			Method:     []byte("INVITE"),
			UriType:    "sip",
			StatusCode: []byte(""),
			StatusDesc: []byte(""),
			User:       []byte("0"),
			Host:       []byte("test"),
			Port:       []byte(""),
			UserType:   []byte(""),
			Src:        []byte("INVITE sip:0@test SIP/2.0"),
		},
		From: message.SipFrom{
			UriType:  "sip",
			Name:     []byte("test"),
			User:     []byte("test"),
			Host:     []byte("test"),
			Port:     []byte("test"),
			Tag:      []byte("e788ae25565649ee9ea69e5fc12345"),
			UserType: []byte(""),
			Src:      []byte(`\"test\" <sip:test@test>;tag=5627f2c27f324ac28fe114c0ef620404"`),
		},
		To:       message.SipTo{},
		Contact:  message.SipContact{},
		Via:      testSipVia,
		Cseq:     message.SipCseq{},
		Auth:     message.SipAuth{},
		Ua:       message.SipVal{},
		Exp:      message.SipVal{},
		MaxFwd:   message.SipVal{},
		CallID:   message.SipVal{},
		ContType: message.SipVal{},
		ContLen:  message.SipVal{},
		Src:      []byte(""),
		Body:     []byte(""),
		Sdp:      sdp.SdpMsg{},
	}
	_ = testSipMsg2
)

func TestSipEvent_GetSipMessage(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	tests := []struct {
		name    string
		fields  fields
		want    *message.SipMsg
		success bool
	}{
		{"Test 1 - ", fields{
			sipmsg: &testSipMsg1,
		}, &testSipMsg1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			if got := se.GetSipMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SipEvent.GetSipMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipEvent_SetSipMessage(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	type args struct {
		msg *message.SipMsg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		success bool
	}{
		// {"Test 1 - ", fields{
		// 	sipmsg: &testSipMsg1,
		// }, args{msg: &testSipMsg1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			fmt.Printf("%#v", tt.args.msg)
			se.SetSipMessage(tt.args.msg)
			fmt.Printf("%#v", tt.args.msg)
			fmt.Println(se.sipmsg)
			fmt.Println(testSipMsg1)
			if !reflect.DeepEqual(se.sipmsg, testSipMsg1) {
				t.Errorf("SipEvent.SetSipMessage() = %v, want %v", se.sipmsg, testSipMsg1)
			}
		})
	}
}

func TestSipEvent_GetTransaction(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	tests := []struct {
		name   string
		fields fields
		want   interfaces.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			if got := se.GetTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SipEvent.GetTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipEvent_SetTransaction(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	type args struct {
		tx interfaces.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			se.SetTransaction(tt.args.tx)
		})
	}
}

func TestSipEvent_SetListeningPoint(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	type args struct {
		lp interfaces.ListeningPoint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			se.SetListeningPoint(tt.args.lp)
		})
	}
}

func TestSipEvent_GetListeningPoint(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	tests := []struct {
		name   string
		fields fields
		want   interfaces.ListeningPoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			if got := se.GetListeningPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SipEvent.GetListeningPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
