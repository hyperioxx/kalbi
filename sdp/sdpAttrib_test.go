package sdp

import "testing"

func Test_sdpAttrib_String(t *testing.T) {
	type fields struct {
		Cat []byte
		Val []byte
		Src []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"sdpAttribTest", fields{Cat: testInBytes, Val: testInBytes, Src: testInBytes}, "a=test:test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := &sdpAttrib{
				Cat: tt.fields.Cat,
				Val: tt.fields.Val,
				Src: tt.fields.Src,
			}
			if got := sa.String(); got != tt.want {
				t.Errorf("sdpAttrib.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSdpAttrib(t *testing.T) {
	type args struct {
		v   []byte
		out *sdpAttrib
	}
	tests := []struct {
		name string
		args args
	}{
		{"sdpAttribTest", args{v: testInBytes, out: &sdpAttrib{Cat: testInBytes, Val: testInBytes, Src: testInBytes}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseSdpAttrib(tt.args.v, tt.args.out)
		})
	}
}
