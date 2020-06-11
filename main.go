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
import "flag"
import "Kalbi/sip/message"
import "Kalbi/transport"
import "Kalbi/dispatcher"
import "Kalbi/sip/transaction"

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

	port := flag.Int("port", 5060, "port number the listening point binds to. default port number 5060")
	host := flag.String("host", "127.0.0.1", "host the listening point binds to. default 127.0.0.1")
	flag.Parse()

	c := make(chan message.Request)
	mainloop := new(dispatcher.EventDispatcher)
	mainloop.AddChannel(c)
	udp := transport.NewTransportListenPoint("udp", *host, *port)
	transactionLayer := transaction.NewManager(c)
	transactionLayer.SetChannel(c)
	go transactionLayer.Start()
    mainloop.AddListenPoint(udp)
    mainloop.Start()
}
