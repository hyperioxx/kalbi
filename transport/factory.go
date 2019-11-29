package transport

import "Kalbi/log"

// NewTransportListenPoint creates liste
func NewTransportListenPoint(protocol string, host string, port int) ListeningPoint {
	switch protocol {
	case "udp":
		log.Log.Info("Creating UDP listening point")
		listner := new(UDPTransport)
		listner.Build(host, port)
		return listner
	case "tcp":
		log.Log.Info("Creating TCP listening point")
		listner := new(UDPTransport)
		listner.Build(host, port)
		return listner
	case "tls":
		log.Log.Info("Creating TLS listening point")
		listner := new(UDPTransport)
		listner.Build(host, port)
		return listner
	default:
		log.Log.Info("Unkown protocol specified")
		panic("Unkown protocol specified")

	}

}


