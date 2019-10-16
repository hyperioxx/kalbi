package main

import server "GoSIP/server"

	

func main(){
	s := server.SIPServer{Host: "127.0.0.1" ,Port:5060}
	s.Listen()
}