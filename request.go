package kalbi

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	method, rest, ok1 := strings.Cut(line, " ")
	requestURI, proto, ok2 := strings.Cut(rest, " ")
	if !ok1 || !ok2 {
		return "", "", "", false
	}
	return method, requestURI, proto, true
}

type Header map[string][]string

func (h Header) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

func (h Header) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}

func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}

func (h Header) Values(key string) []string {
	return textproto.MIMEHeader(h).Values(key)
}

func (h Header) get(key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (h Header) has(key string) bool {
	_, ok := h[key]
	return ok
}

func (h Header) Del(key string) {
	textproto.MIMEHeader(h).Del(key)
}

func (h Header) Clone() Header {
	if h == nil {
		return nil
	}

	nv := 0
	for _, vv := range h {
		nv += len(vv)
	}
	sv := make([]string, nv)
	h2 := make(Header, len(h))
	for k, vv := range h {
		if vv == nil {
			h2[k] = nil
			continue
		}
		n := copy(sv, vv)
		h2[k] = sv[:n:n]
		sv = sv[n:]
	}
	return h2
}

type Request struct {
	Method     string
	Uri        string
	Proto      string
	RemoteAddr *net.UDPAddr
	Header     map[string][]string
}

func readRequest(buffer []byte) (*Request, error) {
	reader := textproto.NewReader(bufio.NewReader(bytes.NewReader(buffer)))
	requestLine, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Error reading the request line:", err)
		return nil, err
	}
	

	method, rest, _ := strings.Cut(requestLine, " ")
	uri, proto, _ := strings.Cut(rest, " ")

	header, err := reader.ReadMIMEHeader()
	if err != nil {
		fmt.Println("Error reading message:", err)
		return nil, err
	}

	return &Request{Method: method , Header: header, Uri: uri, Proto: proto}, nil
}
