package transport

import "Kalbi/log"

// NewTransportListenPoint creates liste
func NewTransportListenPoint(protocol string, host string, port int) ListeningPoint {
	switch protocol {
	case "udp":
		log.Log.Info("Creating UDP listening point")
        
	case "tcp":
		log.Log.Info("Creating TCP listening point")
	default:
		log.Log.Info("Unkown protocol specified")
		panic("Unkown protocol specified")
	}

}