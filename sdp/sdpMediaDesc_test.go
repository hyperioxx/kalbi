package sdp

import "testing"

func Test_sdpMediaDesc_String(t *testing.T) {
	type fields struct {
		MediaType []byte
		Port      []byte
		Proto     []byte
		Fmt       []byte
		Src       []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test sdpMediaDesc_String", fields{MediaType: testInBytes, Port: testInBytes, Proto: testInBytes, Fmt: testInBytes, Src: testInBytes}, "m=test test test test "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &sdpMediaDesc{
				MediaType: tt.fields.MediaType,
				Port:      tt.fields.Port,
				Proto:     tt.fields.Proto,
				Fmt:       tt.fields.Fmt,
				Src:       tt.fields.Src,
			}
			if got := sm.String(); got != tt.want {
				t.Errorf("sdpMediaDesc.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSdpMediaDesc(t *testing.T) {
	type args struct {
		v   []byte
		out *sdpMediaDesc
	}
	tests := []struct {
		name string
		args args
	}{
		{"test parseSdpMediaDesc", args{v: testInBytes, out: &sdpMediaDesc{MediaType: testInBytes, Port: testInBytes, Proto: testInBytes, Fmt: testInBytes, Src: testInBytes}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSdpMediaDesc(tt.args.v, tt.args.out)
			// TODO : add check
		})
	}
}
