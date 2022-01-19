package message

import (
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	type args struct {
		request *SipReq
		via     *SipVia
		to      *SipTo
		from    *SipFrom
		callID  *SipVal
		maxfor  *SipVal
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
			if got := NewResponse(tt.args.request, tt.args.via, tt.args.to, tt.args.from, tt.args.callID, tt.args.maxfor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
