package dialog

import "testing"

func TestGenerateDialogId(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDialogId(); got != tt.want {
				t.Errorf("GenerateDialogId() = %v, want %v", got, tt.want)
			}
		})
	}
}
