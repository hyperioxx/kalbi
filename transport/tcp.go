package transport

import (
	"os"
	"github.com/KalbiProject/Kalbi/log"
	"github.com/KalbiProject/Kalbi/sip/message"
	"net"
)

//TCPTransport is a network protocol listening point for the EventDispatcher
type TCPTransport struct {
    listener *net.TCPListener

}


func (tt *TCPTransport) Read()  *message.SipMsg {

	buffer := make([]byte, 2048)
	conn, err := tt.listener.Accept()
    if err != nil {
        log.Log.Error(err)
	}
	n,  err := conn.Read(buffer)
	
	request := message.Parse(buffer[:n])
	return &request

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
