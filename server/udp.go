package server

import (
	"kalbi/sip/message"
	"net"
)

func NewSipUdpServer(host string, port int) (Server, error) {

	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

	listener, err := net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		return nil, err
	}
	return &SipUdpServer{listener: listener}, nil
}

// SipUdpServer is a network protocol listening point for the EventDispatcher
type SipUdpServer struct {
	listener *net.TCPListener
}

// //Read from TCP Socket
func (tt *SipUdpServer) Read() *message.SipMsg {

	buffer := make([]byte, 2048)
	conn, err := tt.listener.Accept()
	if err != nil {
	}
	n, err := conn.Read(buffer)
	if err != nil {

	}

	request := message.Parse(buffer[:n])
	return &request

}

// Start starts the ListeningPoint
func (tt *SipUdpServer) ListenAndServe() {
}
