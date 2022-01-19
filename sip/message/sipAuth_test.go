package message

import "testing"

func TestSipAuth_SetUsername(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetUsername(tt.args.value)
		})
	}
}

func TestSipAuth_GetUsername(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetUsername(); got != tt.want {
				t.Errorf("SipAuth.GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetRealm(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetRealm(tt.args.value)
		})
	}
}

func TestSipAuth_GetRealm(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetRealm(); got != tt.want {
				t.Errorf("SipAuth.GetRealm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetNonce(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetNonce(tt.args.value)
		})
	}
}

func TestSipAuth_GetNonce(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetNonce(); got != tt.want {
				t.Errorf("SipAuth.GetNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetCNonce(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetCNonce(tt.args.value)
		})
	}
}

func TestSipAuth_GetCNonce(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetCNonce(); got != tt.want {
				t.Errorf("SipAuth.GetCNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetQoP(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetQoP(tt.args.value)
		})
	}
}

func TestSipAuth_GetQoP(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetQoP(); got != tt.want {
				t.Errorf("SipAuth.GetQoP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetAlgorithm(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetAlgorithm(tt.args.value)
		})
	}
}

func TestSipAuth_GetAlgorithm(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetAlgorithm(); got != tt.want {
				t.Errorf("SipAuth.GetAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetNc(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetNc(tt.args.value)
		})
	}
}

func TestSipAuth_GetNc(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetNc(); got != tt.want {
				t.Errorf("SipAuth.GetNc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetURI(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetURI(tt.args.value)
		})
	}
}

func TestSipAuth_GetURI(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetURI(); got != tt.want {
				t.Errorf("SipAuth.GetURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_SetResponse(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			sa.SetResponse(tt.args.value)
		})
	}
}

func TestSipAuth_GetResponse(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.GetResponse(); got != tt.want {
				t.Errorf("SipAuth.GetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipAuth_String(t *testing.T) {
	type fields struct {
		Username  []byte
		Realm     []byte
		Nonce     []byte
		CNonce    []byte
		QoP       []byte
		Algorithm []byte
		Nc        []byte
		URI       []byte
		Response  []byte
		Src       []byte
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
			sa := &SipAuth{
				Username:  tt.fields.Username,
				Realm:     tt.fields.Realm,
				Nonce:     tt.fields.Nonce,
				CNonce:    tt.fields.CNonce,
				QoP:       tt.fields.QoP,
				Algorithm: tt.fields.Algorithm,
				Nc:        tt.fields.Nc,
				URI:       tt.fields.URI,
				Response:  tt.fields.Response,
				Src:       tt.fields.Src,
			}
			if got := sa.String(); got != tt.want {
				t.Errorf("SipAuth.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSipAuth(t *testing.T) {
	type args struct {
		v   []byte
		out *SipAuth
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseSipAuth(tt.args.v, tt.args.out)
		})
	}
}
