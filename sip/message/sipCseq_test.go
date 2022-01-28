package message

import "testing"

func TestSipCseq_SetID(t *testing.T) {
	type fields struct {
		ID     []byte
		Method []byte
		Src    []byte
	}
	type args struct {
		id string
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
			sc := &SipCseq{
				ID:     tt.fields.ID,
				Method: tt.fields.Method,
				Src:    tt.fields.Src,
			}
			sc.SetID(tt.args.id)
		})
	}
}

func TestSipCseq_SetMethod(t *testing.T) {
	type fields struct {
		ID     []byte
		Method []byte
		Src    []byte
	}
	type args struct {
		method string
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
			sc := &SipCseq{
				ID:     tt.fields.ID,
				Method: tt.fields.Method,
				Src:    tt.fields.Src,
			}
			sc.SetMethod(tt.args.method)
		})
	}
}

func TestSipCseq_String(t *testing.T) {
	type fields struct {
		ID     []byte
		Method []byte
		Src    []byte
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
			sc := &SipCseq{
				ID:     tt.fields.ID,
				Method: tt.fields.Method,
				Src:    tt.fields.Src,
			}
			if got := sc.String(); got != tt.want {
				t.Errorf("SipCseq.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSipCseq(t *testing.T) {
	type args struct {
		v   []byte
		out *SipCseq
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipCseq(tt.args.v, tt.args.out)
		})
	}
}
