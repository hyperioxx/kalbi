package sdp

import "testing"

func Test_sdpTime_String(t *testing.T) {
	type fields struct {
		TimeStart []byte
		TimeStop  []byte
		Src       []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test sdpTime_String", fields{TimeStart: testInBytes, TimeStop: testInBytes, Src: testInBytes}, "t=test test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &sdpTime{
				TimeStart: tt.fields.TimeStart,
				TimeStop:  tt.fields.TimeStop,
				Src:       tt.fields.Src,
			}
			if got := st.String(); got != tt.want {
				t.Errorf("sdpTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParserSdpTime(t *testing.T) {
	type args struct {
		v   []byte
		out *sdpTime
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParserSdpTime(tt.args.v, tt.args.out)
			// TODO: Check outputs
		})
	}
}
