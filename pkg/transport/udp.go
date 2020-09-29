package transport

import (
	"Kalbi/pkg/log"
	"Kalbi/pkg/sip/parser"
	"fmt"
	"github.com/marv2097/siprocket"
	"net"
)

//UDPTransport is a network protocol listening point for the EventDispatcher
type UDPTransport struct {
	Address    net.UDPAddr
	Connection *net.UDPConn
}

//Read from UDP Socket
func (ut *UDPTransport) Read() *siprocket.SipMsg {
	buffer := make([]byte, 2048)
	n, _, err := ut.Connection.ReadFromUDP(buffer)
	if err != nil {
		log.Log.Error(err)
	}
	request := parser.Read(buffer[:n])
	return request
}

func (ut *UDPTransport) Build(host string, port int) {
	ut.Address = net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

	var err error
	ut.Connection, err = net.ListenUDP("udp", &ut.Address)
	if err != nil {
		panic(err)
	}

}

func Send(host string, port string, msg string) {
	fmt.Println(msg)
	addr, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		log.Log.Error(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, msg)
	conn.Close()

}
