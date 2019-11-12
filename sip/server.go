package sip

import ("net"
		"fmt"
		"bytes"
		"bufio"
	    )
	

//ListenAndServe function is an endless loop for listening on the specified host and port
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
		n, _, err := conn.ReadFromUDP(buffer)

		if err != nil{
			fmt.Println("GoSIP Error: %v", err)
	    }

		
		bytesbuffer := bytes.NewBuffer(buffer[:n])

		newreader := bufio.NewReader(bytesbuffer)

		request := ReadSIPRequest(newreader)
		function := GetMethodFunction(request.Method)
		function()

		if err != nil{
			panic(err)
		}
	
	}
    

}
