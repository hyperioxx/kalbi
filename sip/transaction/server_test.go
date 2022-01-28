package transaction

import (
	"reflect"
	"testing"

	"github.com/KalbiProject/kalbi/interfaces"
	"github.com/KalbiProject/kalbi/sip/message"
	"github.com/looplab/fsm"
)

func TestServerTransaction_InitFSM(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.InitFSM(tt.args.msg)
		})
	}
}

func TestServerTransaction_SetListeningPoint(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.SetListeningPoint(tt.args.lp)
		})
	}
}

func TestServerTransaction_GetListeningPoint(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			if got := st.GetListeningPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerTransaction.GetListeningPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerTransaction_GetBranchID(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			if got := st.GetBranchID(); got != tt.want {
				t.Errorf("ServerTransaction.GetBranchID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerTransaction_GetOrigin(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			if got := st.GetOrigin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerTransaction.GetOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerTransaction_GetLastMessage(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			if got := st.GetLastMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerTransaction.GetLastMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerTransaction_SetLastMessage(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.SetLastMessage(tt.args.msg)
		})
	}
}

func TestServerTransaction_Receive(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.Receive(tt.args.msg)
		})
	}
}

func TestServerTransaction_Respond(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.Respond(tt.args.msg)
		})
	}
}

func TestServerTransaction_GetServerTransactionID(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			if got := st.GetServerTransactionID(); got != tt.want {
				t.Errorf("ServerTransaction.GetServerTransactionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServerTransaction_Send(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.Send(tt.args.msg, tt.args.host, tt.args.port)
		})
	}
}

func TestServerTransaction_actRespond(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.actRespond(tt.args.event)
		})
	}
}

func TestServerTransaction_actRespondDelete(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.actRespondDelete(tt.args.event)
		})
	}
}

func TestServerTransaction_actTransErr(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.actTransErr(tt.args.event)
		})
	}
}

func TestServerTransaction_actDelete(t *testing.T) {
	type fields struct {
		ID             string
		BranchID       string
		TransManager   *TransactionManager
		Origin         *message.SipMsg
		FSM            *fsm.FSM
		msgHistory     []*message.SipMsg
		ListeningPoint interfaces.ListeningPoint
		Host           string
		Port           string
		LastMessage    *message.SipMsg
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
			st := &ServerTransaction{
				ID:             tt.fields.ID,
				BranchID:       tt.fields.BranchID,
				TransManager:   tt.fields.TransManager,
				Origin:         tt.fields.Origin,
				FSM:            tt.fields.FSM,
				msgHistory:     tt.fields.msgHistory,
				ListeningPoint: tt.fields.ListeningPoint,
				Host:           tt.fields.Host,
				Port:           tt.fields.Port,
				LastMessage:    tt.fields.LastMessage,
			}
			st.actDelete(tt.args.event)
		})
	}
}
