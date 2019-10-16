package server

import ("net"
		"fmt")


	

type SIPServer struct {
	Host string
	Port int
}


//Use the HandleRequest() to set functions for each Method type e.g. INVITE, REGISTRATION, INFO, CANCEL
func (s *SIPServer) HandleRequest(method string, ){
	
}

//Lisen() function is an endless loop for listening on the specified host and port
func (s *SIPServer) Listen(){
	fmt.Println("Starting Go Sip Server")
	fmt.Printf("Listening on %s:%d\n", s.Host, s.Port)

	buffer := make([]byte, 2048)

	addr := net.UDPAddr{
		                IP : net.ParseIP(s.Host),
		                Port: s.Port,
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