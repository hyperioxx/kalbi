package kalbi

import (
	"context"
	"fmt"
	"net"
	"time"
)

type UDPServer struct {
	Addr    string
	Handler Handler
	Conn    *net.UDPConn
}

func NewUDPServer(addr string, handler Handler) SIPServer {
	return &UDPServer{
		Addr:    addr,
		Handler: handler,
	}
}

func (s *UDPServer) ListenAndServe() error {
	addr, err := net.ResolveUDPAddr("udp", s.Addr)
	if err != nil {
		return err
	}

	s.Conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer s.Conn.Close()
	fmt.Println("SIP UDP Server started on", s.Addr)

	return s.serve()
}

func (s *UDPServer) serve() error {
	for {
		select {
		case <-context.Background().Done():
			return nil
		default:
			buffer := make([]byte, 2048)
			n, remoteAddr, err := s.Conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}
			req, err := readRequest(buffer[:n])

			req.RemoteAddr = remoteAddr

			responseWriter := &UDPResponseWriter{conn: s.Conn, addr: remoteAddr}

			go s.Handler.ServeSIP(responseWriter, req)
		}
	}
}

func (s *UDPServer) Shutdown(ctx context.Context) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(5 * time.Second)
	}
	s.Conn.SetReadDeadline(deadline)
	<-ctx.Done()
	return s.Conn.Close()
}

type UDPResponseWriter struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

func (rw *UDPResponseWriter) Write(message []byte) (int, error) {
	return rw.conn.WriteToUDP(message, rw.addr)
}

func (rw *UDPResponseWriter) RemoteAddr() net.Addr {
	return rw.addr
}
