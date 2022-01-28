package sdp

import (
	"bytes"
	"testing"
)

var (
	connDataString = []byte{116, 101, 115, 116, 110, 101, 116, 32, 116, 101, 115, 116, 97, 100, 100, 114, 32, 116, 101, 115, 116, 99, 111, 110, 110, 32, 116, 101, 115, 116, 115, 114, 99}
	netTypeOut     = []byte{116, 101, 115, 116, 110, 101, 116}
	addrTypeOut    = []byte{116, 101, 115, 116, 97, 100, 100, 114}
	connAddrOut    = []byte{116, 101, 115, 116, 99, 111, 110, 110}
	srcOut         = []byte{116, 101, 115, 116, 110, 101, 116, 32, 116, 101, 115, 116, 97, 100, 100, 114, 32, 116, 101, 115, 116, 99, 111, 110, 110, 32, 116, 101, 115, 116, 115, 114, 99}
)

func Test_sdpConnData_String(t *testing.T) {
	type fields struct {
		NetType  []byte
		AddrType []byte
		ConnAddr []byte
		Src      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test sdpConnData String", fields{NetType: testInBytes, AddrType: testInBytes, ConnAddr: testInBytes, Src: testInBytes}, "c=test test test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &sdpConnData{
				NetType:  tt.fields.NetType,
				AddrType: tt.fields.AddrType,
				ConnAddr: tt.fields.ConnAddr,
				Src:      tt.fields.Src,
			}
			if got := sc.String(); got != tt.want {
				t.Errorf("sdpConnData.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSdpConnectionData(t *testing.T) {
	type args struct {
		v   []byte
		out *sdpConnData
	}
	tests := []struct {
		name string
		args args
		want *sdpConnData
	}{
		{"test 1 parseSdpConnectionDate whitespace trigger",
			args{v: connDataString, out: &sdpConnData{}},
			&sdpConnData{NetType: netTypeOut, AddrType: addrTypeOut, ConnAddr: connAddrOut, Src: srcOut}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSdpConnectionData(tt.args.v, tt.args.out)
			if !bytes.Equal(tt.args.out.NetType, tt.want.NetType) {
				t.Errorf("parseSdpConnectionData() NetType = %v, want %v", tt.args.out.NetType, tt.want.NetType)
			}
			if !bytes.Equal(tt.args.out.AddrType, tt.want.AddrType) {
				t.Errorf("parseSdpConnectionData() AddrType = %v, want %v", tt.args.out.AddrType, tt.want.AddrType)
			}
			if !bytes.Equal(tt.args.out.ConnAddr, tt.want.ConnAddr) {
				t.Errorf("parseSdpConnectionData() ConnAddr = %v, want %v", tt.args.out.ConnAddr, tt.want.ConnAddr)
			}
			if !bytes.Equal(tt.args.out.Src, tt.want.Src) {
				t.Errorf("parseSdpConnectionData() Src = %v, want %v", tt.args.out.Src, tt.want.Src)
			}
		})
	}
}
