package server

import ("net"
		"fmt")
		


func SIPServer(host string , port int){
	fmt.Println("Starting Go Sip Server")
	fmt.Printf("Listening on %s:%d\n", host, port)

	buffer := make([]byte, 2048)

	addr := net.UDPAddr{
		                IP : net.ParseIP(host),
		                Port: port,
					   }
					   
	conn, err := net.ListenUDP("udp", &addr)

	if err != nil{
		fmt.Println("GoSIP Error: %v", err)
	}
	for {
		_, remoteaddr, err := conn.ReadFromUDP(buffer)

		if err != nil{
			fmt.Println("GoSIP Error: %v", err)
		}

		fmt.Printf("Message: %v %s ", remoteaddr, buffer)
	}
    

}