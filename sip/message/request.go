package message

import "bufio"
import "net/textproto"
import "fmt"
import "strings"
import "io"
import "net/url"


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
	Via *url.URL
	Contact *url.URL
	To *url.URL
	From *url.URL
	Branch string
}


