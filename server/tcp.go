package server

import (
	"kalbi/sip/message"
	"net"
)

func NewSipTCPServer(host string, port int) (Server, error) {

	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

	listener, err := net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		return nil, err
	}
	return &SipTCPServer{listener: listener}, nil
}

type SipTCPServer struct {
	listener *net.TCPListener
}

func (tt *SipTCPServer) Read() *message.SipMsg {

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
func (tt *SipTCPServer) ListenAndServe() {
}
