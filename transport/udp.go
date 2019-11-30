package transport

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"Kalbi/sip/message"
	//"Kalbi/log"
	//"Kalbi/sip/transaction"
)



//UDPTransport is a network protocol listening point for the EventDispatcher
type UDPTransport struct {
    Address net.UDPAddr
	Connection *net.UDPConn
}

//Read from UDP Socket
func (ut *UDPTransport) Read() *message.Request{
	buffer := make([]byte, 2048)
	n, _, err := ut.Connection.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
	}
	bytesbuffer := bytes.NewBuffer(buffer[:n])
	newreader := bufio.NewReader(bytesbuffer)
	request := message.Read(newreader)
    return request
}


func (ut *UDPTransport) Build(host string, port int){
	ut.Address = net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}

	var err error
	ut.Connection, err  = net.ListenUDP("udp", &ut.Address)
	if err != nil{
		panic(err)
	}

}



/*
//ListenAndServe function is an endless loop for listening on the specified host and port
func ListenAndServe(Host string, Port int) {
    log.Log.Info("Starting Kalbi Server")
	log.Log.Info("Listening on "+ Host)
	
    inputChannel := make(chan *message.Request)

	transManager := transaction.NewManager(inputChannel)

	go transManager.Start()
	
	buffer := make([]byte, 2048)

	addr := net.UDPAddr{
		IP:   net.ParseIP(Host),
		Port: Port,
	}

	conn , err = net.ListenUDP("udp", &addr)

	if err != nil {
		log.Log.Error("Unable to bind to socket")
	}

	for {
		n, _, err := conn.ReadFromUDP(buffer)

		if err != nil {
			log.Log.Error("Unable to read from socket")
		}
        
		bytesbuffer := bytes.NewBuffer(buffer[:n])

		newreader := bufio.NewReader(bytesbuffer)

		request := message.Read(newreader)
		
		inputChannel <- request

		if err != nil {
			panic(err)
		}

	}

}
*/