package message

import (
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	type args struct {
		statuscode int
		tx         Transaction
	}
	tests := []struct {
		name string
		args args
		want *SipMsg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(tt.args.tx, tt.args.statuscode, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
