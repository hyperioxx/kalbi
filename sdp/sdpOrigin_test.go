package sdp

import (
	"bytes"
	"testing"
)

var (
	testSdpOriginStr = "testusername testsessionid testsessionversion testnet testaddr testuni testsrc "
)

func TestSdpOrigin_String(t *testing.T) {
	type fields struct {
		Username       []byte
		SessionID      []byte
		SessionVersion []byte
		NetType        []byte
		AddrType       []byte
		UniAddr        []byte
		Src            []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test TestSdpOrigin_String", fields{Username: testInBytes, SessionID: testInBytes, SessionVersion: testInBytes, NetType: testInBytes, AddrType: testInBytes, UniAddr: testInBytes, Src: testInBytes}, "o=test test test test test test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			so := &SdpOrigin{
				Username:       tt.fields.Username,
				SessionID:      tt.fields.SessionID,
				SessionVersion: tt.fields.SessionVersion,
				NetType:        tt.fields.NetType,
				AddrType:       tt.fields.AddrType,
				UniAddr:        tt.fields.UniAddr,
				Src:            tt.fields.Src,
			}
			if got := so.String(); got != tt.want {
				t.Errorf("SdpOrigin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSdpOrigin(t *testing.T) {
	type args struct {
		v   []byte
		out *SdpOrigin
	}
	tests := []struct {
		name string
		args args
		want *SdpOrigin
	}{
		{"Test 1", args{
			v:   []byte(testSdpOriginStr),
			out: &SdpOrigin{}},
			&SdpOrigin{
				Username:       []byte("testusername"),
				SessionID:      []byte("testsessionid"),
				SessionVersion: []byte("testsessionversion"),
				NetType:        []byte("testnet"),
				AddrType:       []byte("testaddr"),
				UniAddr:        []byte("testuni"),
				Src:            []byte("testusername testsessionid testsessionversion testnet testaddr testuni testsrc "),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSdpOrigin(tt.args.v, tt.args.out)
			if !bytes.Equal(tt.args.out.Username, tt.want.Username) {
				t.Errorf("parseSdpConnectionData() Username = %v, want %v", tt.args.out.Username, tt.want.Username)
			}
			if !bytes.Equal(tt.args.out.SessionID, tt.want.SessionID) {
				t.Errorf("parseSdpConnectionData() SessionID = %v, want %v", tt.args.out.SessionID, tt.want.SessionID)
			}
			if !bytes.Equal(tt.args.out.SessionVersion, tt.want.SessionVersion) {
				t.Errorf("parseSdpConnectionData() SessionVersion = %v, want %v", tt.args.out.SessionVersion, tt.want.SessionVersion)
			}
			if !bytes.Equal(tt.args.out.NetType, tt.want.NetType) {
				t.Errorf("parseSdpConnectionData() NetType = %v, want %v", tt.args.out.NetType, tt.want.NetType)
			}
			if !bytes.Equal(tt.args.out.AddrType, tt.want.AddrType) {
				t.Errorf("parseSdpConnectionData() AddrType = %v, want %v", tt.args.out.AddrType, tt.want.AddrType)
			}
			if !bytes.Equal(tt.args.out.UniAddr, tt.want.UniAddr) {
				t.Errorf("parseSdpConnectionData() ConnAddr = %v, want %v", tt.args.out.UniAddr, tt.want.UniAddr)
			}
			if !bytes.Equal(tt.args.out.Src, tt.want.Src) {
				t.Errorf("parseSdpConnectionData() Src = %v, want %v", tt.args.out.Src, tt.want.Src)
			}
		})
	}
}
