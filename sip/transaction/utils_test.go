package transaction

import "testing"

func TestGenerateBranchD(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateBranchID(); got != tt.want {
				t.Errorf("GenerateBranchID() = %v, want %v", got, tt.want)
			}
		})
	}
}
