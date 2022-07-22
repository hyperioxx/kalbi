package logger

import (
	"errors"
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestGetLogger(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: set up tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLogger()
			if gotStr := fmt.Sprintf("%#v", got); gotStr != tt.want {
				t.Errorf("GetLogger() = %v, want %v", gotStr, tt.want)
			}

		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		msg  string
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.tags...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		msg  string
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg, tt.args.tags...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg  string
		err  error
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", err: errors.New("test"), tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}

func TestCritical(t *testing.T) {
	type args struct {
		msg  string
		err  error
		tags []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test 1", args{msg: "test", err: errors.New("test"), tags: []zap.Field{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Critical(tt.args.msg, tt.args.err, tt.args.tags...)
		})
	}
}
