package main

import server "GoSIP/server"

	

func main(){
	server.SIPServer("127.0.0.1", 5060)
}