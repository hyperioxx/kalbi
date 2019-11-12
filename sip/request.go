package sip

import "bufio"
import "net/textproto"
import "fmt"
import "strings"
import "io"


type SIPRequest struct{
	Method string
	RequestURI string
	Proto string
	Headers map[string]string
}



// ReadSIPRequest reads and parses an icoming request
func ReadSIPRequest(br *bufio.Reader) *SIPRequest{
	reader := newTextProtoReader(br)
	requestline, err := reader.ReadLine()

	var req = new(SIPRequest)

	if err != nil {
	  fmt.Println(err)
	}

	req.Method , req.RequestURI , req.Proto = parseRequestLine(requestline)
	req.Headers = parseHeaders(reader)

    return req
	  
}


func newTextProtoReader(br *bufio.Reader) *textproto.Reader {
	return textproto.NewReader(br)
}


func parseRequestLine(requestline string) (method string , uri string, proto string) {
	parsedline := strings.Split(requestline[1:], " ")
	return parsedline[0] , parsedline[1] , parsedline[2]
}


func parseHeaders(reader *textproto.Reader)  map[string]string {

	var headers map[string]string

	for{
		line, err := reader.ReadLine()
		if err == io.EOF{
		   break
		}
		if err != nil{
			fmt.Println(err)
		}
		if line != "<nil>" {
			headermatch := strings.Index(line,":")
			if headermatch != -1 {
				headers[strings.ToLower(line[:headermatch])] = line[headermatch + 1:]
			}
		}
	 }
	
	return headers
}