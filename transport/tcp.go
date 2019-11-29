package transport


import "net"

//TCPTransport is a network protocol listening point for the EventDispatcher
type TCPTransport struct {

}


func (tt *TCPTransport) Read(){

}


func (ut *TCPTransport) Build(host string, port int){
	ut.Buffer = make([]byte, 2048)
	ut.Address = net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

}