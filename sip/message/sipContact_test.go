package message

import "testing"

func TestSipContact_SetUriType(t *testing.T) {
	type fields struct {
		UriType string
		Name    []byte
		User    []byte
		Host    []byte
		Port    []byte
		Tran    []byte
		Qval    []byte
		Expires []byte
		Src     []byte
	}
	type args struct {
		uriType string
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
			sc := &SipContact{
				UriType: tt.fields.UriType,
				Name:    tt.fields.Name,
				User:    tt.fields.User,
				Host:    tt.fields.Host,
				Port:    tt.fields.Port,
				Tran:    tt.fields.Tran,
				Qval:    tt.fields.Qval,
				Expires: tt.fields.Expires,
				Src:     tt.fields.Src,
			}
			sc.SetUriType(tt.args.uriType)
		})
	}
}

func TestSipContact_SetUser(t *testing.T) {
	type fields struct {
		UriType string
		Name    []byte
		User    []byte
		Host    []byte
		Port    []byte
		Tran    []byte
		Qval    []byte
		Expires []byte
		Src     []byte
	}
	type args struct {
		user string
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
			sc := &SipContact{
				UriType: tt.fields.UriType,
				Name:    tt.fields.Name,
				User:    tt.fields.User,
				Host:    tt.fields.Host,
				Port:    tt.fields.Port,
				Tran:    tt.fields.Tran,
				Qval:    tt.fields.Qval,
				Expires: tt.fields.Expires,
				Src:     tt.fields.Src,
			}
			sc.SetUser(tt.args.user)
		})
	}
}

func TestSipContact_SetHost(t *testing.T) {
	type fields struct {
		UriType string
		Name    []byte
		User    []byte
		Host    []byte
		Port    []byte
		Tran    []byte
		Qval    []byte
		Expires []byte
		Src     []byte
	}
	type args struct {
		host string
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
			sc := &SipContact{
				UriType: tt.fields.UriType,
				Name:    tt.fields.Name,
				User:    tt.fields.User,
				Host:    tt.fields.Host,
				Port:    tt.fields.Port,
				Tran:    tt.fields.Tran,
				Qval:    tt.fields.Qval,
				Expires: tt.fields.Expires,
				Src:     tt.fields.Src,
			}
			sc.SetHost(tt.args.host)
		})
	}
}

func TestSipContact_SetPort(t *testing.T) {
	type fields struct {
		UriType string
		Name    []byte
		User    []byte
		Host    []byte
		Port    []byte
		Tran    []byte
		Qval    []byte
		Expires []byte
		Src     []byte
	}
	type args struct {
		port string
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
			sc := &SipContact{
				UriType: tt.fields.UriType,
				Name:    tt.fields.Name,
				User:    tt.fields.User,
				Host:    tt.fields.Host,
				Port:    tt.fields.Port,
				Tran:    tt.fields.Tran,
				Qval:    tt.fields.Qval,
				Expires: tt.fields.Expires,
				Src:     tt.fields.Src,
			}
			sc.SetPort(tt.args.port)
		})
	}
}

func TestSipContact_SetName(t *testing.T) {
	type fields struct {
		UriType string
		Name    []byte
		User    []byte
		Host    []byte
		Port    []byte
		Tran    []byte
		Qval    []byte
		Expires []byte
		Src     []byte
	}
	type args struct {
		name string
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
			sc := &SipContact{
				UriType: tt.fields.UriType,
				Name:    tt.fields.Name,
				User:    tt.fields.User,
				Host:    tt.fields.Host,
				Port:    tt.fields.Port,
				Tran:    tt.fields.Tran,
				Qval:    tt.fields.Qval,
				Expires: tt.fields.Expires,
				Src:     tt.fields.Src,
			}
			sc.SetName(tt.args.name)
		})
	}
}

func TestSipContact_String(t *testing.T) {
	type fields struct {
		UriType string
		Name    []byte
		User    []byte
		Host    []byte
		Port    []byte
		Tran    []byte
		Qval    []byte
		Expires []byte
		Src     []byte
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
			sc := &SipContact{
				UriType: tt.fields.UriType,
				Name:    tt.fields.Name,
				User:    tt.fields.User,
				Host:    tt.fields.Host,
				Port:    tt.fields.Port,
				Tran:    tt.fields.Tran,
				Qval:    tt.fields.Qval,
				Expires: tt.fields.Expires,
				Src:     tt.fields.Src,
			}
			if got := sc.String(); got != tt.want {
				t.Errorf("SipContact.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSipContact(t *testing.T) {
	type args struct {
		v   []byte
		out *SipContact
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipContact(tt.args.v, tt.args.out)
		})
	}
}
