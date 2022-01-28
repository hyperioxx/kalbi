package transport

//import (
// "net"
// "reflect"
// "testing"

// "github.com/KalbiProject/kalbi/sip/message"
//)

// func TestTCPTransport_Read(t *testing.T) {
// 	type fields struct {
// 		listener         *net.TCPListener
// 		TransportChannel chan *message.SipMsg
// 		connTable        map[string]net.Conn
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   *message.SipMsg
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt := &TCPTransport{
// 				listener:         tt.fields.listener,
// 				TransportChannel: tt.fields.TransportChannel,
// 				connTable:        tt.fields.connTable,
// 			}
// 			if got := tt.Read(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TCPTransport.Read() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTCPTransport_Start(t *testing.T) {
// 	type fields struct {
// 		listener         *net.TCPListener
// 		TransportChannel chan *message.SipMsg
// 		connTable        map[string]net.Conn
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt := &TCPTransport{
// 				listener:         tt.fields.listener,
// 				TransportChannel: tt.fields.TransportChannel,
// 				connTable:        tt.fields.connTable,
// 			}
// 			tt.Start()
// 		})
// 	}
// }

// func TestTCPTransport_SetTransportChannel(t *testing.T) {
// 	type fields struct {
// 		listener         *net.TCPListener
// 		TransportChannel chan *message.SipMsg
// 		connTable        map[string]net.Conn
// 	}
// 	type args struct {
// 		channel chan *message.SipMsg
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt := &TCPTransport{
// 				listener:         tt.fields.listener,
// 				TransportChannel: tt.fields.TransportChannel,
// 				connTable:        tt.fields.connTable,
// 			}
// 			tt.SetTransportChannel(tt.args.channel)
// 		})
// 	}
// }

// func TestTCPTransport_Build(t *testing.T) {
// 	type fields struct {
// 		listener         *net.TCPListener
// 		TransportChannel chan *message.SipMsg
// 		connTable        map[string]net.Conn
// 	}
// 	type args struct {
// 		host string
// 		port int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt := &TCPTransport{
// 				listener:         tt.fields.listener,
// 				TransportChannel: tt.fields.TransportChannel,
// 				connTable:        tt.fields.connTable,
// 			}
// 			tt.Build(tt.args.host, tt.args.port)
// 		})
// 	}
// }
