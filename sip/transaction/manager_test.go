package transaction

import (
	"reflect"
	"sync"
	"testing"

	"github.com/KalbiProject/kalbi/interfaces"
	"github.com/KalbiProject/kalbi/sip/message"
)

func TestNewTransactionManager(t *testing.T) {
	tests := []struct {
		name string
		want *TransactionManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionManager_Handle(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		event interfaces.SipEventObject
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interfaces.SipEventObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			if got := tm.Handle(tt.args.event); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionManager_FindServerTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		msg *message.SipMsg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interfaces.Transaction
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			got, got1 := tm.FindServerTransaction(tt.args.msg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.FindServerTransaction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TransactionManager.FindServerTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTransactionManager_FindClientTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		msg *message.SipMsg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interfaces.Transaction
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			got, got1 := tm.FindClientTransaction(tt.args.msg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.FindClientTransaction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TransactionManager.FindClientTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTransactionManager_FindServerTransactionByID(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interfaces.Transaction
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			got, got1 := tm.FindServerTransactionByID(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.FindServerTransactionByID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TransactionManager.FindServerTransactionByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTransactionManager_FindClientTransactionByID(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interfaces.Transaction
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			got, got1 := tm.FindClientTransactionByID(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.FindClientTransactionByID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TransactionManager.FindClientTransactionByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTransactionManager_PutServerTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
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
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			tm.PutServerTransaction(tt.args.tx)
		})
	}
}

func TestTransactionManager_PutClientTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
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
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			tm.PutClientTransaction(tt.args.tx)
		})
	}
}

func TestTransactionManager_DeleteServerTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		branch string
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
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			tm.DeleteServerTransaction(tt.args.branch)
		})
	}
}

func TestTransactionManager_DeleteClientTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		branch string
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
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			tm.DeleteClientTransaction(tt.args.branch)
		})
	}
}

func TestTransactionManager_MakeKey(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		msg message.SipMsg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			if got := tm.MakeKey(tt.args.msg); got != tt.want {
				t.Errorf("TransactionManager.MakeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionManager_NewClientTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		msg *message.SipMsg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ClientTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			if got := tm.NewClientTransaction(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.NewClientTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionManager_NewServerTransaction(t *testing.T) {
	type fields struct {
		ServerTX        map[string]interfaces.Transaction
		ClientTX        map[string]interfaces.Transaction
		RequestChannel  chan interfaces.Transaction
		ResponseChannel chan interfaces.Transaction
		ListeningPoint  interfaces.ListeningPoint
		txLock          *sync.RWMutex
	}
	type args struct {
		msg *message.SipMsg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ServerTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TransactionManager{
				ServerTX:        tt.fields.ServerTX,
				ClientTX:        tt.fields.ClientTX,
				RequestChannel:  tt.fields.RequestChannel,
				ResponseChannel: tt.fields.ResponseChannel,
				ListeningPoint:  tt.fields.ListeningPoint,
				txLock:          tt.fields.txLock,
			}
			if got := tm.NewServerTransaction(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionManager.NewServerTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
