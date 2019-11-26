package parser

import "bufio"
import "net/textproto"
import "fmt"
import "strings"
import "io"
import "net/url"
import "Kalbi/sip/message"


func Read(br *bufio.Reader) *message.Request {
	reader := newTextProtoReader(br)
	requestline, err := reader.ReadLine()

	var req = new(message.Request)

	if err != nil {
		fmt.Println(err)
	}

	req.Method, req.RequestURI, req.Proto = parseRequestLine(requestline)
	req.Headers = parseHeaders(reader)

	/*req.Via , err = parseVia(req.Headers["via"])
	if err != nil {
		fmt.Println(err)
	}*/
	req.Contact, err = parseURI(req.Headers["contact"])
	if err != nil {
		fmt.Println(err)
	}

	req.From, err = parseURI(req.Headers["from"])
	if err != nil {
		fmt.Println(err)
	}

	req.To, err = parseURI(req.Headers["to"])
	if err != nil {
		fmt.Println(err)
	}

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



func parseURI(_uri string) (*url.URL, error) {
	_uri = cleanURI(_uri)
	uri, err := url.Parse(_uri)
	return uri, err
}


func cleanURI(uri string) string{
	uri = uri[1:len(uri)-1]
	return uri
}