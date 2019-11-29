package main
/* 
			██╗  ██╗ █████╗ ██╗     ██████╗ ██╗ 
			██║ ██╔╝██╔══██╗██║     ██╔══██╗██║
			█████╔╝ ███████║██║     ██████╔╝██║
			██╔═██╗ ██╔══██║██║     ██╔══██╗██║
			██║  ██╗██║  ██║███████╗██████╔╝██║
			╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝
			===================================
			SIP Application Server in GoLang
			Maintainer: Aaron Parfitt
*/



import "fmt"

//import "Kalbi/log"
import "Kalbi/transport"

var title string = `
			██╗  ██╗ █████╗ ██╗     ██████╗ ██╗ 
			██║ ██╔╝██╔══██╗██║     ██╔══██╗██║
			█████╔╝ ███████║██║     ██████╔╝██║
			██╔═██╗ ██╔══██║██║     ██╔══██╗██║
			██║  ██╗██║  ██║███████╗██████╔╝██║
			╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═════╝ ╚═╝
			===================================
			 SIP Application Server in GoLang.
			===================================`

func main() {
	fmt.Println(title)
	transport.NewTransportListenPoint("udp", "127.0.0.1", 5060)
	
	//transport.ListenAndServe("127.0.0.1", 5060)
}
