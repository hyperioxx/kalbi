package dialog

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var (
	dlgID = int32(1213456)
	dlg   = Dialog{
		DialogID: dlgID,
		CallID:   "testCallerID",
		ToTag:    "testToTag",
		FromTag:  "testFromTag",
	}
	dlgMapNil = make(map[string]Dialog)
	dlgMapdlg = map[string]Dialog{}
	rw        = sync.RWMutex{}
)

func TestNewDialogManager(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test 1 - NewDialogManager", "*dialog.DialogManager"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDialogManager(); fmt.Sprintf("%T", got) != "*dialog.DialogManager" {
				t.Errorf("NewDialogManager() = %v, want %v", fmt.Sprintf("%T", got), tt.want)
			}
		})
	}
}

func TestDialogManager_GetDialog(t *testing.T) {

	dlgMapdlg["test"] = dlg

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
		{"Test 1 - NewDialogManager - nil", fields{dialogs: dlgMapNil, Lock: &rw}, args{value: ""}, nil},
		{"Test 2 - NewDialogManager - basic dialog", fields{dialogs: dlgMapdlg, Lock: &rw}, args{value: "test"}, &dlg},
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
		{"Test 1 - TestDialogManager_DeleteDialog - nil", fields{dialogs: dlgMapNil, Lock: &rw}, args{value: ""}},
		{"Test 2 - TestDialogManager_DeleteDialog - basic dialog", fields{dialogs: dlgMapdlg, Lock: &rw}, args{value: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &DialogManager{
				dialogs: tt.fields.dialogs,
				Lock:    tt.fields.Lock,
			}
			dm.DeleteDialog(tt.args.value)
			if dm.dialogs[tt.args.value].DialogID != 0 || dm.dialogs[tt.args.value].CallID != "" || dm.dialogs[tt.args.value].ToTag != "" || dm.dialogs[tt.args.value].FromTag != "" {
				t.Errorf(`DialogManager.GetDialog() = %v, want 0 for int32 or "" for string fields`, dm.dialogs[tt.args.value])
			}

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
	}{
		{"Test 1 - TestDialogManager_NewDialog - nil", fields{dialogs: dlgMapNil, Lock: &rw}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &DialogManager{
				dialogs: tt.fields.dialogs,
				Lock:    tt.fields.Lock,
			}
			got := dm.NewDialog()
			if reflect.TypeOf(got.DialogID).Kind() != reflect.Int32 {
				t.Errorf("GenerateDialogID() = %v, want %v", fmt.Sprintf("%T", got), "int32")
			}
			if got.CallID != "" || got.ToTag != "" || got.FromTag != "" {
				t.Error("GenerateDialogID() = string key is not blank")

			}
			if got.ClientTx != nil || got.ServerTx != nil {
				t.Error("GenerateDialogID() = interface key is not blank")
			}
		})
	}
}
