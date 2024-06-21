package kalbi

import (
	"context"
	"net"
)

type SIPServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type ResponseWriter interface {
	Write(message []byte) (int, error)
	RemoteAddr() net.Addr
}

type Handler interface {
	ServeSIP(ctx ResponseWriter, req *Request)
}
