package transport

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"io"
	"strings"
	"Kalbi/sip/parser"
	"Kalbi/log"
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
	//bytesbuffer := bytes.NewBuffer(buffer[:n])
	//newreader := bufio.NewReader(bytesbuffer)
	//buf := new(strings.Builder)
	//_, err = io.Copy(buf, newreader)
	//if err != nil {
		log.Log.Error(err)
	//}

	request := message.Read(string(buffer[:n]))
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


func UdpSend(host string, port string, msg string){
	addr, err := net.ResolveUDPAddr("udp", host + ":" + port)
	if err != nil{
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

