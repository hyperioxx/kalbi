package sdp

import "testing"

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
		{"test sdpConnData String",fields{NetType: testInBytes,AddrType: testInBytes, ConnAddr: testInBytes, Src: testInBytes },"c=test test test"},
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
	}{
		{"test parseSdpConnectionDate",args{v:testInBytes,out: &sdpConnData{NetType: testInBytes,AddrType: testInBytes,ConnAddr: testInBytes,Src: testInBytes}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSdpConnectionData(tt.args.v, tt.args.out)
			// TODO : add check 
		})
	}
}
