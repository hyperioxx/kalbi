package message

import "bufio"
import "net/textproto"
import "fmt"
import "strings"
import "io"



func Read(br *bufio.Reader) *Request {
	reader := newTextProtoReader(br)
	requestline, err := reader.ReadLine()

	var req = new(Request)

	if err != nil {
		fmt.Println(err)
	}

	req.Method, req.RequestURI, req.Proto = parseRequestLine(requestline)
	req.Headers = parseHeaders(reader)

	req.Contact = parseURI(req.Headers["contact"])
	req.From = parseURI(req.Headers["from"])
	req.To = parseURI(req.Headers["to"])
	//req.Via = parseURI(req.Headers["via"])
	return req

}


func newTextProtoReader(br *bufio.Reader) *textproto.Reader {
	return textproto.NewReader(br)
}


func parseRequestLine(requestline string) (method string, uri string, proto string) {
	parsedline := strings.Split(requestline[1:], " ")
	return parsedline[0], parsedline[1], parsedline[2]
}

func parseHeaders(reader *textproto.Reader) map[string]string {

	headers := make(map[string]string)

	for {
		line, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if line != "<nil>" {
			headermatch := strings.Index(line, ":")
			if headermatch != -1 {
				headers[strings.ToLower(line[:headermatch])] = line[headermatch+1:]
			}
		}
	}

	return headers
}


