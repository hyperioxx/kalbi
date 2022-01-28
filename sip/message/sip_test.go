package message

import (
	"reflect"
	"testing"

	"github.com/KalbiProject/kalbi/sdp"
)

func TestSipMsg_SetAuthHeader(t *testing.T) {
	type fields struct {
		Req      SipReq
		From     SipFrom
		To       SipTo
		Contact  SipContact
		Via      []SipVia
		Cseq     SipCseq
		Auth     SipAuth
		Ua       SipVal
		Exp      SipVal
		MaxFwd   SipVal
		CallID   SipVal
		ContType SipVal
		ContLen  SipVal
		Src      []byte
		Body     []byte
		Sdp      sdp.SdpMsg
	}
	type args struct {
		auth *SipAuth
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
			sm := &SipMsg{
				Req:      tt.fields.Req,
				From:     tt.fields.From,
				To:       tt.fields.To,
				Contact:  tt.fields.Contact,
				Via:      tt.fields.Via,
				Cseq:     tt.fields.Cseq,
				Auth:     tt.fields.Auth,
				Ua:       tt.fields.Ua,
				Exp:      tt.fields.Exp,
				MaxFwd:   tt.fields.MaxFwd,
				CallID:   tt.fields.CallID,
				ContType: tt.fields.ContType,
				ContLen:  tt.fields.ContLen,
				Src:      tt.fields.Src,
				Body:     tt.fields.Body,
				Sdp:      tt.fields.Sdp,
			}
			sm.SetAuthHeader(tt.args.auth)
		})
	}
}

func TestSipMsg_GetStatusCode(t *testing.T) {
	type fields struct {
		Req      SipReq
		From     SipFrom
		To       SipTo
		Contact  SipContact
		Via      []SipVia
		Cseq     SipCseq
		Auth     SipAuth
		Ua       SipVal
		Exp      SipVal
		MaxFwd   SipVal
		CallID   SipVal
		ContType SipVal
		ContLen  SipVal
		Src      []byte
		Body     []byte
		Sdp      sdp.SdpMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SipMsg{
				Req:      tt.fields.Req,
				From:     tt.fields.From,
				To:       tt.fields.To,
				Contact:  tt.fields.Contact,
				Via:      tt.fields.Via,
				Cseq:     tt.fields.Cseq,
				Auth:     tt.fields.Auth,
				Ua:       tt.fields.Ua,
				Exp:      tt.fields.Exp,
				MaxFwd:   tt.fields.MaxFwd,
				CallID:   tt.fields.CallID,
				ContType: tt.fields.ContType,
				ContLen:  tt.fields.ContLen,
				Src:      tt.fields.Src,
				Body:     tt.fields.Body,
				Sdp:      tt.fields.Sdp,
			}
			if got := sm.GetStatusCode(); got != tt.want {
				t.Errorf("SipMsg.GetStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipMsg_CopyHeaders(t *testing.T) {
	type fields struct {
		Req      SipReq
		From     SipFrom
		To       SipTo
		Contact  SipContact
		Via      []SipVia
		Cseq     SipCseq
		Auth     SipAuth
		Ua       SipVal
		Exp      SipVal
		MaxFwd   SipVal
		CallID   SipVal
		ContType SipVal
		ContLen  SipVal
		Src      []byte
		Body     []byte
		Sdp      sdp.SdpMsg
	}
	type args struct {
		msg *SipMsg
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
			sm := &SipMsg{
				Req:      tt.fields.Req,
				From:     tt.fields.From,
				To:       tt.fields.To,
				Contact:  tt.fields.Contact,
				Via:      tt.fields.Via,
				Cseq:     tt.fields.Cseq,
				Auth:     tt.fields.Auth,
				Ua:       tt.fields.Ua,
				Exp:      tt.fields.Exp,
				MaxFwd:   tt.fields.MaxFwd,
				CallID:   tt.fields.CallID,
				ContType: tt.fields.ContType,
				ContLen:  tt.fields.ContLen,
				Src:      tt.fields.Src,
				Body:     tt.fields.Body,
				Sdp:      tt.fields.Sdp,
			}
			sm.CopyHeaders(tt.args.msg)
		})
	}
}

func TestSipMsg_CopySdp(t *testing.T) {
	type fields struct {
		Req      SipReq
		From     SipFrom
		To       SipTo
		Contact  SipContact
		Via      []SipVia
		Cseq     SipCseq
		Auth     SipAuth
		Ua       SipVal
		Exp      SipVal
		MaxFwd   SipVal
		CallID   SipVal
		ContType SipVal
		ContLen  SipVal
		Src      []byte
		Body     []byte
		Sdp      sdp.SdpMsg
	}
	type args struct {
		msg *SipMsg
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
			sm := &SipMsg{
				Req:      tt.fields.Req,
				From:     tt.fields.From,
				To:       tt.fields.To,
				Contact:  tt.fields.Contact,
				Via:      tt.fields.Via,
				Cseq:     tt.fields.Cseq,
				Auth:     tt.fields.Auth,
				Ua:       tt.fields.Ua,
				Exp:      tt.fields.Exp,
				MaxFwd:   tt.fields.MaxFwd,
				CallID:   tt.fields.CallID,
				ContType: tt.fields.ContType,
				ContLen:  tt.fields.ContLen,
				Src:      tt.fields.Src,
				Body:     tt.fields.Body,
				Sdp:      tt.fields.Sdp,
			}
			sm.CopySdp(tt.args.msg)
		})
	}
}

func TestSipMsg_String(t *testing.T) {
	type fields struct {
		Req      SipReq
		From     SipFrom
		To       SipTo
		Contact  SipContact
		Via      []SipVia
		Cseq     SipCseq
		Auth     SipAuth
		Ua       SipVal
		Exp      SipVal
		MaxFwd   SipVal
		CallID   SipVal
		ContType SipVal
		ContLen  SipVal
		Src      []byte
		Body     []byte
		Sdp      sdp.SdpMsg
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SipMsg{
				Req:      tt.fields.Req,
				From:     tt.fields.From,
				To:       tt.fields.To,
				Contact:  tt.fields.Contact,
				Via:      tt.fields.Via,
				Cseq:     tt.fields.Cseq,
				Auth:     tt.fields.Auth,
				Ua:       tt.fields.Ua,
				Exp:      tt.fields.Exp,
				MaxFwd:   tt.fields.MaxFwd,
				CallID:   tt.fields.CallID,
				ContType: tt.fields.ContType,
				ContLen:  tt.fields.ContLen,
				Src:      tt.fields.Src,
				Body:     tt.fields.Body,
				Sdp:      tt.fields.Sdp,
			}
			if got := sm.String(); got != tt.want {
				t.Errorf("SipMsg.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipVal_SetValue(t *testing.T) {
	type fields struct {
		Value []byte
		Src   []byte
	}
	type args struct {
		value string
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
			sv := &SipVal{
				Value: tt.fields.Value,
				Src:   tt.fields.Src,
			}
			sv.SetValue(tt.args.value)
		})
	}
}

func TestSipVal_String(t *testing.T) {
	type fields struct {
		Value []byte
		Src   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := &SipVal{
				Value: tt.fields.Value,
				Src:   tt.fields.Src,
			}
			if got := sv.String(); got != tt.want {
				t.Errorf("SipVal.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		v []byte
	}
	tests := []struct {
		name       string
		args       args
		wantOutput SipMsg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Parse(tt.args.v); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Parse() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_indexSep(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := indexSep(tt.args.s)
			if got != tt.want {
				t.Errorf("indexSep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("indexSep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getString(t *testing.T) {
	type args struct {
		sl   []byte
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getString(tt.args.sl, tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("getString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBytes(t *testing.T) {
	type args struct {
		sl   []byte
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBytes(tt.args.sl, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageDetails(t *testing.T) {
	type args struct {
		data *SipMsg
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MessageDetails(tt.args.data); got != tt.want {
				t.Errorf("MessageDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
