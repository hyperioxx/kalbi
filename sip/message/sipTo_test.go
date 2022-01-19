package message

import "testing"

func TestSipTo_SetUriType(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			sf.SetUriType(tt.args.uriType)
		})
	}
}

func TestSipTo_SetUser(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			sf.SetUser(tt.args.user)
		})
	}
}

func TestSipTo_SetHost(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			sf.SetHost(tt.args.host)
		})
	}
}

func TestSipTo_SetPort(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			sf.SetPort(tt.args.port)
		})
	}
}

func TestSipTo_SetUserType(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			sf.SetUserType(tt.args.userType)
		})
	}
}

func TestSipTo_SetTag(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
	}
	type args struct {
		tag string
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			sf.SetTag(tt.args.tag)
		})
	}
}

func TestSipTo_String(t *testing.T) {
	type fields struct {
		UriType  string
		Name     []byte
		User     []byte
		Host     []byte
		Port     []byte
		Tag      []byte
		UserType []byte
		Src      []byte
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
			sf := &SipTo{
				UriType:  tt.fields.UriType,
				Name:     tt.fields.Name,
				User:     tt.fields.User,
				Host:     tt.fields.Host,
				Port:     tt.fields.Port,
				Tag:      tt.fields.Tag,
				UserType: tt.fields.UserType,
				Src:      tt.fields.Src,
			}
			if got := sf.String(); got != tt.want {
				t.Errorf("SipTo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSipTo(t *testing.T) {
	type args struct {
		v   []byte
		out *SipTo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipTo(tt.args.v, tt.args.out)
		})
	}
}
