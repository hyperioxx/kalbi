package parser

import( "bufio"
        "net/textproto"
        "fmt"
        "strings"
        "io"
        "bytes"
        "regexp"
        "Kalbi/sip/message"
        "github.com/pixelbender/go-sdp/sdp")




type Parser struct {
	
}





func Read(msg string) *message.Request {




	

	return req

}




func newTextProtoReader(br *bufio.Reader) *textproto.Reader {
	return textproto.NewReader(br)
}


func parseRequestLine(requestline string) (method string, uri string, proto string) {
	parsedline := strings.Split(requestline[0:], " ")
	return parsedline[0], parsedline[1], parsedline[2]
}

func parseHeaders(reader *textproto.Reader) textproto.MIMEHeader {

	mimeHeader, err := reader.ReadMIMEHeader()

	
	if err != nil{
		fmt.Println(err)
	}

	return mimeHeader
}


func parseSDP(body string) *sdp.Session{


	d := strings.Split(body, "\r\n\r\n")

	session, err := sdp.ParseString(d[1])
	if err != nil{
		fmt.Println(err)
	}
   
	fmt.Println(session)
	return session

}


func parseVia(via string) string{
	if via == "" {
		return ""
	}
	re := regexp.MustCompile(`(branch=([^;].*))`)
	match := re.FindStringSubmatch(via)
	return match[2]
}

