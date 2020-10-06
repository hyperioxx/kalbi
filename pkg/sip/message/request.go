package message


func NewRequest(request string) *Request{
	r := new(Request)
	return r

}

type Request struct {
	ResponseCode string
	RequestURI   string
	Proto        string
	Headers      map[string]string
	Via          []*URI
	Contact      []*URI
	To           *URI
	From         *URI
	Branch       string
	Cseq         string
}

func (r *Request) CopyHeader(){

}