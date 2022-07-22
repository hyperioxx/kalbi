package sdp

import "testing"

func Test_sdpVersion_String(t *testing.T) {
	type fields struct {
		Val []byte
		Src []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test sdpVersionString", fields{Val: testInBytes, Src: testInBytes}, "v=test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := &sdpVersion{
				Val: tt.fields.Val,
				Src: tt.fields.Src,
			}
			if got := sv.String(); got != tt.want {
				t.Errorf("sdpVersion.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
