package message

import (
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	type args struct {
		request *SipReq
		via     *SipVia
		to      *SipTo
		from    *SipFrom
		contact *SipContact
		callID  *SipVal
		cseq    *SipCseq
		maxfor  *SipVal
		contlen *SipVal
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
			if got := NewRequest(tt.args.request, tt.args.via, tt.args.to, tt.args.from, tt.args.contact, tt.args.callID, tt.args.cseq, tt.args.maxfor, tt.args.contlen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
