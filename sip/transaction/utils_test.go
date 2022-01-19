package transaction

import "testing"

func TestGenerateBranchId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateBranchId(); got != tt.want {
				t.Errorf("GenerateBranchId() = %v, want %v", got, tt.want)
			}
		})
	}
}
