package transport

import (
	"github.com/KalbiProject/kalbi/sip/message"
	"reflect"
	"testing"
)

type ListenPoint struct {
	Want message.ListeningPoint
}

func TestNewTransportListenPoint(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     int
	}
	output := ListenPoint{}
	_ = output // remove when tests updated
	tests := []struct {
		name string
		args args
		want message.ListeningPoint
	}{
		// {"test 1", args{protocol: "test", host: "test", port: 5000}, output.Want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransportListenPoint(tt.args.protocol, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransportListenPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
