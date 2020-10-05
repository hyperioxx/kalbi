package transport

import (
	"os"
	"Kalbi/pkg/log"
	"Kalbi/pkg/sip/parser"
	"github.com/marv2097/siprocket"
	"net"
)

//TCPTransport is a network protocol listening point for the EventDispatcher
type TCPTransport struct {
    listener *net.TCPListener

}


func (tt *TCPTransport) Read()  *siprocket.SipMsg {

	buffer := make([]byte, 2048)
	conn, err := tt.listener.Accept()
    if err != nil {
        log.Log.Error(err)
	}
	n,  err := conn.Read(buffer)
	
	request := parser.Read(buffer[:n])
	return request

}


func (tt *TCPTransport) Build(host string, port int){
	var err error
	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}
	
	tt.listener, err = net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		log.Log.Error(err)
		os.Exit(1)
	}


}
