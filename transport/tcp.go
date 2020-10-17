package transport

import (
	"github.com/KalbiProject/Kalbi/log"
	"github.com/KalbiProject/Kalbi/sip/message"
	"net"
	"os"
)

//TCPTransport is a network protocol listening point for the EventDispatcher
type TCPTransport struct {
	listener         *net.TCPListener
	TransportChannel chan *message.SipMsg
	connTable        map[string] net.Conn
}

func (tt *TCPTransport) Read() *message.SipMsg {

	buffer := make([]byte, 2048)
	conn, err := tt.listener.Accept()
	tt.connTable[conn.RemoteAddr().String()] = conn
	if err != nil {
		log.Log.Error(err)
	}
	n, err := conn.Read(buffer)
	if err != nil {
		log.Log.Error(err)
	}

	request := message.Parse(buffer[:n])
	return &request

}

func (tt *TCPTransport) Start() {
	log.Log.Info("Starting TCP Listening Point ")
	for {
		msg := tt.Read()
		tt.TransportChannel <-msg
	}
}

func (tt *TCPTransport) SetTransportChannel(channel chan *message.SipMsg) {
	tt.TransportChannel = channel
}


func (tt *TCPTransport) Build(host string, port int) {
	var err error
	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

	tt.connTable = make(map[string] net.Conn)
	tt.listener, err = net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		log.Log.Error(err)
		os.Exit(1)
	}

}
