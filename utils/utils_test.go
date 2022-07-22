package utils

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		s   []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test Contains 1 - true", args{s: []string{"a", "b", "c", "d"}, str: "a"}, true},
		{"test Contains 2 - false", args{s: []string{"a", "b", "c", "d"}, str: "e"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.str, tt.args.s); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsByte(t *testing.T) {
	type args struct {
		b  []byte
		by byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test Contains 1 - true", args{b: []byte{'a', 'b', 'c', 'd'}, by: 'a'}, true},
		{"test Contains 2 - false", args{b: []byte{'a', 'b', 'c', 'd'}, by: 'e'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsByte(tt.args.by, tt.args.b); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
