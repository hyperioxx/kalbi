package sdp

import "testing"

func TestSdpOrigin_String(t *testing.T) {
	type fields struct {
		Username       []byte
		SessionId      []byte
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
		{"test TestSdpOrigin_String", fields{Username: testInBytes, SessionId: testInBytes, SessionVersion: testInBytes, NetType: testInBytes, AddrType: testInBytes, UniAddr: testInBytes,Src: testInBytes}, "o=test test test test test test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			so := &SdpOrigin{
				Username:       tt.fields.Username,
				SessionId:      tt.fields.SessionId,
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSdpOrigin(tt.args.v, tt.args.out)
			// TODO: check input/outputs
		})
	}
}
