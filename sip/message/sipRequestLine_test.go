package message

import "testing"

func TestSipReq_SetMethod(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetMethod(tt.args.method)
		})
	}
}

func TestSipReq_SetUriType(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetUriType(tt.args.uriType)
		})
	}
}

func TestSipReq_SetStatusCode(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
	}
	type args struct {
		code int
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetStatusCode(tt.args.code)
		})
	}
}

func TestSipReq_SetStatusDesc(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
	}
	type args struct {
		desc string
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetStatusDesc(tt.args.desc)
		})
	}
}

func TestSipReq_SetUser(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetUser(tt.args.user)
		})
	}
}

func TestSipReq_SetHost(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetHost(tt.args.host)
		})
	}
}

func TestSipReq_SetPort(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetPort(tt.args.port)
		})
	}
}

func TestSipReq_SetUserType(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
	}
	type args struct {
		userType string
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			sr.SetUserType(tt.args.userType)
		})
	}
}

func TestSipReq_String(t *testing.T) {
	type fields struct {
		Method     []byte
		UriType    string
		StatusCode []byte
		StatusDesc []byte
		User       []byte
		Host       []byte
		Port       []byte
		UserType   []byte
		Src        []byte
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
			sr := &SipReq{
				Method:     tt.fields.Method,
				UriType:    tt.fields.UriType,
				StatusCode: tt.fields.StatusCode,
				StatusDesc: tt.fields.StatusDesc,
				User:       tt.fields.User,
				Host:       tt.fields.Host,
				Port:       tt.fields.Port,
				UserType:   tt.fields.UserType,
				Src:        tt.fields.Src,
			}
			if got := sr.String(); got != tt.want {
				t.Errorf("SipReq.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSipReq(t *testing.T) {
	type args struct {
		v   []byte
		out *SipReq
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipReq(tt.args.v, tt.args.out)
		})
	}
}
