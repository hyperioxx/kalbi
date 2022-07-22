package transaction

import (
	"reflect"
	"testing"
	"time"

	"github.com/KalbiProject/kalbi/interfaces"
	"github.com/KalbiProject/kalbi/sip/message"
	"github.com/looplab/fsm"
)

func TestClientTransaction_InitFSM(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.InitFSM(tt.args.msg)
		})
	}
}

func TestClientTransaction_SetListeningPoint(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.SetListeningPoint(tt.args.lp)
		})
	}
}

func TestClientTransaction_GetListeningPoint(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			if got := ct.GetListeningPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientTransaction.GetListeningPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientTransaction_GetBranchID(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			if got := ct.GetBranchID(); got != tt.want {
				t.Errorf("ClientTransaction.GetBranchID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientTransaction_GetOrigin(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			if got := ct.GetOrigin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientTransaction.GetOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientTransaction_Receive(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.Receive(tt.args.msg)
		})
	}
}

func TestClientTransaction_act100(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		event *fsm.Event
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.act100(tt.args.event)
		})
	}
}

func TestClientTransaction_SetServerTransaction(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		txID string
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.SetServerTransaction(tt.args.txID)
		})
	}
}

func TestClientTransaction_GetServerTransactionID(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			if got := ct.GetServerTransactionID(); got != tt.want {
				t.Errorf("ClientTransaction.GetServerTransactionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientTransaction_GetLastMessage(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			if got := ct.GetLastMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientTransaction.GetLastMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientTransaction_SetLastMessage(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.SetLastMessage(tt.args.msg)
		})
	}
}

func TestClientTransaction_actSend(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		event *fsm.Event
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.actSend(tt.args.event)
		})
	}
}

func TestClientTransaction_act300(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		event *fsm.Event
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.act300(tt.args.event)
		})
	}
}

func TestClientTransaction_actTransErr(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		event *fsm.Event
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.actTransErr(tt.args.event)
		})
	}
}

func TestClientTransaction_actDelete(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		event *fsm.Event
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.actDelete(tt.args.event)
		})
	}
}

func TestClientTransaction_actResend(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		event *fsm.Event
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.actResend(tt.args.event)
		})
	}
}

func TestClientTransaction_Resend(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.Resend()
		})
	}
}

func TestClientTransaction_StatelessSend(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		msg  *message.SipMsg
		host string
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.StatelessSend(tt.args.msg, tt.args.host, tt.args.port)
		})
	}
}

func TestClientTransaction_Send(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		ServerTxID     string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
		timerATime     time.Duration
		timerA         *time.Timer
		timerB         *time.Timer
		timerDTime     time.Duration
		timerD         *time.Timer
	}
	type args struct {
		msg  *message.SipMsg
		host string
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
			ct := &ClientTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				ServerTxID:     tt.fields.ServerTxID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
				timerATime:     tt.fields.timerATime,
				timerA:         tt.fields.timerA,
				timerB:         tt.fields.timerB,
				timerDTime:     tt.fields.timerDTime,
				timerD:         tt.fields.timerD,
			}
			ct.Send(tt.args.msg, tt.args.host, tt.args.port)
		})
	}
}
