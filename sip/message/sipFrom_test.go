package message

import "testing"

func TestSipFrom_SetUriType(t *testing.T) {
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
			sf := &SipFrom{
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

func TestSipFrom_SetUser(t *testing.T) {
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
			sf := &SipFrom{
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

func TestSipFrom_SetHost(t *testing.T) {
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
			sf := &SipFrom{
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

func TestSipFrom_SetPort(t *testing.T) {
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
			sf := &SipFrom{
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

func TestSipFrom_SetUserType(t *testing.T) {
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
			sf := &SipFrom{
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

func TestSipFrom_SetTag(t *testing.T) {
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
			sf := &SipFrom{
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

func TestSipFrom_String(t *testing.T) {
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
			sf := &SipFrom{
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
				t.Errorf("SipFrom.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSipFrom(t *testing.T) {
	type args struct {
		v   []byte
		out *SipFrom
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipFrom(tt.args.v, tt.args.out)
		})
	}
}
