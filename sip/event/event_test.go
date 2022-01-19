package event

import (
	"reflect"
	"testing"

	"github.com/KalbiProject/Kalbi/interfaces"
	"github.com/KalbiProject/Kalbi/sip/message"
)

func TestSipEvent_GetSipMessage(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	tests := []struct {
		name   string
		fields fields
		want   *message.SipMsg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			if got := se.GetSipMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SipEvent.GetSipMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipEvent_SetSipMessage(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	type args struct {
		msg *message.SipMsg
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
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			se.SetSipMessage(tt.args.msg)
		})
	}
}

func TestSipEvent_GetTransaction(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	tests := []struct {
		name   string
		fields fields
		want   interfaces.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			if got := se.GetTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SipEvent.GetTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSipEvent_SetTransaction(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	type args struct {
		tx interfaces.Transaction
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
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			se.SetTransaction(tt.args.tx)
		})
	}
}

func TestSipEvent_SetListeningPoint(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	type args struct {
		lp interfaces.ListeningPoint
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
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			se.SetListeningPoint(tt.args.lp)
		})
	}
}

func TestSipEvent_GetListeningPoint(t *testing.T) {
	type fields struct {
		sipmsg *message.SipMsg
		tx     interfaces.Transaction
		lp     interfaces.ListeningPoint
	}
	tests := []struct {
		name   string
		fields fields
		want   interfaces.ListeningPoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := &SipEvent{
				sipmsg: tt.fields.sipmsg,
				tx:     tt.fields.tx,
				lp:     tt.fields.lp,
			}
			if got := se.GetListeningPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SipEvent.GetListeningPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
