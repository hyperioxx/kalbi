package message

import (
	"reflect"
	"testing"
)

func TestNewRequestLine(t *testing.T) {
	type args struct {
		method  string
		uriType string
		user    string
		host    string
		port    string
	}
	tests := []struct {
		name string
		args args
		want *SipReq
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequestLine(tt.args.method, tt.args.uriType, tt.args.user, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequestLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResponseLine(t *testing.T) {
	type args struct {
		statuscode int
		statusdesc string
	}
	tests := []struct {
		name string
		args args
		want *SipReq
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponseLine(tt.args.statuscode, tt.args.statusdesc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewViaHeader(t *testing.T) {
	type args struct {
		transport string
		host      string
		port      string
	}
	tests := []struct {
		name string
		args args
		want *SipVia
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewViaHeader(tt.args.transport, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewViaHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromHeader(t *testing.T) {
	type args struct {
		user    string
		uriType string
		host    string
		port    string
	}
	tests := []struct {
		name string
		args args
		want *SipFrom
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromHeader(tt.args.user, tt.args.uriType, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewToHeader(t *testing.T) {
	type args struct {
		user    string
		uriType string
		host    string
		port    string
	}
	tests := []struct {
		name string
		args args
		want *SipTo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToHeader(tt.args.user, tt.args.uriType, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContactHeader(t *testing.T) {
	type args struct {
		uriType string
		user    string
		host    string
	}
	tests := []struct {
		name string
		args args
		want *SipContact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContactHeader(tt.args.uriType, tt.args.user, tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContactHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCallID(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *SipVal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCallID(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCallID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCSeq(t *testing.T) {
	type args struct {
		id     string
		method string
	}
	tests := []struct {
		name string
		args args
		want *SipCseq
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSeq(tt.args.id, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMaxForwards(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *SipVal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaxForwards(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaxForwards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContentLength(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *SipVal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContentLength(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContentLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
