package dialog

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewDialogManager(t *testing.T) {
	tests := []struct {
		name string
		want *DialogManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDialogManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDialogManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialogManager_GetDialog(t *testing.T) {
	type fields struct {
		dialogs map[string]Dialog
		Lock    *sync.RWMutex
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dialog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &DialogManager{
				dialogs: tt.fields.dialogs,
				Lock:    tt.fields.Lock,
			}
			if got := dm.GetDialog(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DialogManager.GetDialog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialogManager_DeleteDialog(t *testing.T) {
	type fields struct {
		dialogs map[string]Dialog
		Lock    *sync.RWMutex
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
			dm := &DialogManager{
				dialogs: tt.fields.dialogs,
				Lock:    tt.fields.Lock,
			}
			dm.DeleteDialog(tt.args.value)
		})
	}
}

func TestDialogManager_NewDialog(t *testing.T) {
	type fields struct {
		dialogs map[string]Dialog
		Lock    *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *Dialog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &DialogManager{
				dialogs: tt.fields.dialogs,
				Lock:    tt.fields.Lock,
			}
			if got := dm.NewDialog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DialogManager.NewDialog() = %v, want %v", got, tt.want)
			}
		})
	}
}
