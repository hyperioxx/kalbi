package message


//Request sip request struct
type Request struct {
	Method     string
	RequestURI string
	Proto      string
	Headers    map[string]string
	Via        *URI
	Contact    *URI
	To         *URI
	From       *URI
	Branch     string
}
