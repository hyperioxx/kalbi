package sip

import ("net"
		"fmt"
		"bytes"
		"bufio"
	    "net/textproto")
//

	

//ListenAndServe() function is an endless loop for listening on the specified host and port
func ListenAndServe(Host string, Port int){
	fmt.Println("Starting Go Sip Server")
	fmt.Printf("Listening on %s:%d\n", Host, Port)

	buffer := make([]byte, 2048)

	addr := net.UDPAddr{
		                IP : net.ParseIP(Host),
		                Port: Port,
					   }
					   
	conn, err := net.ListenUDP("udp", &addr)

	if err != nil{
		fmt.Println("GoSIP Error: %v", err)
	}
	for {
		_, remoteaddr, err := conn.ReadFrom(buffer)

		if err != nil{
			fmt.Println("GoSIP Error: %v", err)
		}

		fmt.Println(remoteaddr)
		
		bytesbuffer := bytes.NewBuffer(buffer)

		newreader := bufio.NewReader(bytesbuffer)

		tp := textproto.NewReader(newreader)
        for{
		headers, err := tp.ReadLine()
		if err == "" {
			break
		}
		if err != nil{
			panic(err)
		}
		
		fmt.Println(headers)
	}
	}
    

}
