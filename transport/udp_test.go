package transport

import (
	"net"
	"reflect"
	"testing"

	"github.com/KalbiProject/kalbi/interfaces"
)

func TestUDPTransport_Read(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	tests := []struct {
		name   string
		fields fields
		want   interfaces.SipEventObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			if got := ut.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UDPTransport.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUDPTransport_GetHost(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			if got := ut.GetHost(); got != tt.want {
				t.Errorf("UDPTransport.GetHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUDPTransport_GetPort(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			if got := ut.GetPort(); got != tt.want {
				t.Errorf("UDPTransport.GetPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUDPTransport_Build(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	type args struct {
		host string
		port int
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
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			ut.Build(tt.args.host, tt.args.port)
		})
	}
}

func TestUDPTransport_Start(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			ut.Start()
		})
	}
}

func TestUDPTransport_SetTransportChannel(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	type args struct {
		channel chan interfaces.SipEventObject
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
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			ut.SetTransportChannel(tt.args.channel)
		})
	}
}

func TestUDPTransport_Send(t *testing.T) {
	type fields struct {
		Host             string
		Port             int
		Address          net.UDPAddr
		Connection       net.PacketConn
		TransportChannel chan interfaces.SipEventObject
	}
	type args struct {
		host string
		port string
		msg  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UDPTransport{
				Host:             tt.fields.Host,
				Port:             tt.fields.Port,
				Address:          tt.fields.Address,
				Connection:       tt.fields.Connection,
				TransportChannel: tt.fields.TransportChannel,
			}
			if err := ut.Send(tt.args.host, tt.args.port, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("UDPTransport.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
