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

import (
	"Kalbi/pkg/sip/stack"
	//"Kalbi/internal/kalbi/sip/transaction"
	//"Kalbi/internal/kalbi/transport"
	"flag"
	//"github.com/marv2097/siprocket"
)


func main() {
	

	port := flag.Int("port", 5060, "port number the listening point binds to. default port number 5060")
	host := flag.String("host", "127.0.0.1", "host the listening point binds to. default 127.0.0.1")
	//mode := flag.String("mode", "127.0.0.1", "host the listening point binds to. default 127.0.0.1")
	flag.Parse()

	//c := make(chan siprocket.SipMsg)


	mainloop := new(stack.SipStack)
	//mainloop.AddChannel(c)
	//transactionLayer := transaction.NewManager(c)
	//transactionLayer.SetChannel(c)
	//go transactionLayer.Start()
	mainloop.CreateListenPoint("udp", *host, *port)
	mainloop.Start()
}
