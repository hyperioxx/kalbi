package message

import "bufio"
import "net/textproto"
import "fmt"
import "strings"
import "io"


//RequestContainer is a collection of address and request
type RequestContainer struct {
	address string
	request Request
}

//Request sip request struct
type Request struct {
	Method     string
	RequestURI string
	Proto      string
	Headers    map[string]string
	Branch string
}

// ReadSIPRequest reads and parses an icoming request
func ReadSIPRequest(br *bufio.Reader) *Request {
	reader := newTextProtoReader(br)
	requestline, err := reader.ReadLine()

	var req = new(Request)

	if err != nil {
		fmt.Println(err)
	}

	req.Method, req.RequestURI, req.Proto = parseRequestLine(requestline)
	req.Headers = parseHeaders(reader)
	req.Branch , err = parseBranch(req.Headers["via"])
	
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


func parseBranch(via string) (string, error) {
	parsedLine := strings.Split(via, ";")
	for i := range parsedLine{
		param := strings.Split(parsedLine[i+1], "=")
		if param[0] == "branch"{
			return param[1], nil
		}
	}
return "", nil
}
