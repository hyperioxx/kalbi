package message

import "testing"

func TestSipVia_String(t *testing.T) {
	type fields struct {
		Trans  string
		Host   []byte
		Port   []byte
		Branch []byte
		Rport  []byte
		Maddr  []byte
		Ttl    []byte
		Rcvd   []byte
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
			sv := &SipVia{
				Trans:  tt.fields.Trans,
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
				Branch: tt.fields.Branch,
				Rport:  tt.fields.Rport,
				Maddr:  tt.fields.Maddr,
				Ttl:    tt.fields.Ttl,
				Rcvd:   tt.fields.Rcvd,
				Src:    tt.fields.Src,
			}
			if got := sv.String(); got != tt.want {
				t.Errorf("SipVia.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipVia_SetTransport(t *testing.T) {
	type fields struct {
		Trans  string
		Host   []byte
		Port   []byte
		Branch []byte
		Rport  []byte
		Maddr  []byte
		Ttl    []byte
		Rcvd   []byte
		Src    []byte
	}
	type args struct {
		trans string
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
			sv := &SipVia{
				Trans:  tt.fields.Trans,
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
				Branch: tt.fields.Branch,
				Rport:  tt.fields.Rport,
				Maddr:  tt.fields.Maddr,
				Ttl:    tt.fields.Ttl,
				Rcvd:   tt.fields.Rcvd,
				Src:    tt.fields.Src,
			}
			sv.SetTransport(tt.args.trans)
		})
	}
}

func TestSipVia_SetHost(t *testing.T) {
	type fields struct {
		Trans  string
		Host   []byte
		Port   []byte
		Branch []byte
		Rport  []byte
		Maddr  []byte
		Ttl    []byte
		Rcvd   []byte
		Src    []byte
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
			sv := &SipVia{
				Trans:  tt.fields.Trans,
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
				Branch: tt.fields.Branch,
				Rport:  tt.fields.Rport,
				Maddr:  tt.fields.Maddr,
				Ttl:    tt.fields.Ttl,
				Rcvd:   tt.fields.Rcvd,
				Src:    tt.fields.Src,
			}
			sv.SetHost(tt.args.value)
		})
	}
}

func TestSipVia_SetPort(t *testing.T) {
	type fields struct {
		Trans  string
		Host   []byte
		Port   []byte
		Branch []byte
		Rport  []byte
		Maddr  []byte
		Ttl    []byte
		Rcvd   []byte
		Src    []byte
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
			sv := &SipVia{
				Trans:  tt.fields.Trans,
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
				Branch: tt.fields.Branch,
				Rport:  tt.fields.Rport,
				Maddr:  tt.fields.Maddr,
				Ttl:    tt.fields.Ttl,
				Rcvd:   tt.fields.Rcvd,
				Src:    tt.fields.Src,
			}
			sv.SetPort(tt.args.value)
		})
	}
}

func TestSipVia_SetBranch(t *testing.T) {
	type fields struct {
		Trans  string
		Host   []byte
		Port   []byte
		Branch []byte
		Rport  []byte
		Maddr  []byte
		Ttl    []byte
		Rcvd   []byte
		Src    []byte
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
			sv := &SipVia{
				Trans:  tt.fields.Trans,
				Host:   tt.fields.Host,
				Port:   tt.fields.Port,
				Branch: tt.fields.Branch,
				Rport:  tt.fields.Rport,
				Maddr:  tt.fields.Maddr,
				Ttl:    tt.fields.Ttl,
				Rcvd:   tt.fields.Rcvd,
				Src:    tt.fields.Src,
			}
			sv.SetBranch(tt.args.value)
		})
	}
}

func TestParseSipVia(t *testing.T) {
	type args struct {
		v   []byte
		out *SipVia
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipVia(tt.args.v, tt.args.out)
		})
	}
}
