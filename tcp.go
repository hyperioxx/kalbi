package kalbi

import (
	"context"
	"fmt"
	"io"
	"net"
)

type TCPServer struct {
	Addr     string
	Handler  Handler
	Listener net.Listener
}

func NewTCPServer(addr string, handler Handler) SIPServer {
	return &TCPServer{
		Addr:    addr,
		Handler: handler,
	}
}

func (s *TCPServer) ListenAndServe() error {
	var err error
	s.Listener, err = net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer s.Listener.Close()

	fmt.Println("TCP Server started on", s.Addr)

	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			return err
		}

		// Handle each connection in a new goroutine
		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	responseWriter := &TCPResponseWriter{conn: conn}
	for {

		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err)
			}
			break
		}

		req, err := readRequest(buffer[:n])
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		if s.Handler != nil {
			s.Handler.ServeSIP(responseWriter, req)
		}
	}
}

func (s *TCPServer) Shutdown(ctx context.Context) error {
	if err := s.Listener.Close(); err != nil {
		return err
	}

	<-ctx.Done()

	return ctx.Err()
}

type TCPResponseWriter struct {
	conn net.Conn
}

func (rw *TCPResponseWriter) Write(message []byte) (int, error) {
	return rw.conn.Write(message)
}

func (rw *TCPResponseWriter) RemoteAddr() net.Addr {
	return rw.conn.RemoteAddr()
}
