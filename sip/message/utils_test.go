package message

import "testing"

func TestGenerateBranchID(t *testing.T) {
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

func TestGenerateNewCallID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNewCallID(); got != tt.want {
				t.Errorf("GenerateNewCallID() = %v, want %v", got, tt.want)
			}
		})
	}
}
