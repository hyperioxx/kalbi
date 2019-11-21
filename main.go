package main

import "GoSIP/sip"

	

func main(){
	
	sip.ListenAndServe("127.0.0.1", 5060)
}