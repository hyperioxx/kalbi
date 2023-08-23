package message

// Parses a single line that is in the format of a to line, v
// Also requires a pointer to a struct of type sipTo to write output to
// RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt - 8.1.1.2 To

//SipTo SIP To Header
type SipTo struct {
	UriType  string // Type of URI sip, sips, tel etc
	Name     []byte // Named portion of URI
	User     []byte // User part
	Host     []byte // Host part
	Port     []byte // Port number
	Tag      []byte // Tag
	UserType []byte // User Type
	Src      []byte // Full source if needed
}

// Scan makes it easier to parse messages
type Scan struct {
	body []byte
	pos  int
}

// NewScan is the constructor for Scan
func NewScan(body []byte) *Scan {
	return &Scan{body, 0}
}

// CaptureString will capture matching options and consume the characters, returns captured string and bool
func (scan *Scan) CaptureString(options ...string) (string, bool) {
	for _, option := range options {
		if getString(scan.body, scan.pos, scan.pos+len(option)) == option {
			scan.pos += len(option)
			return option, true
		}
	}
	return "", false
}

// ConsumeString will capture matching options and consume the characters, returns true if one of options found
func (scan *Scan) ConsumeString(options ...string) bool {
	_, found := scan.CaptureString(options...)
	return found
}

// CaptureChar will capture a single character from the options, returns the current char and a boolean (is the character one of the options)
func (scan *Scan) CaptureChar(options ...byte) (byte, bool) {
	currentChar := scan.body[scan.pos]
	for _, option := range options {
		if option == currentChar {
			scan.pos++
			return currentChar, true
		}
	}
	return currentChar, false
}

// CheckChar looks for matches ahead but does not consume
func (scan *Scan) CheckChar(options ...byte) bool {
	currentChar := scan.body[scan.pos]
	for _, option := range options {
		if option == currentChar {
			return true
		}
	}
	return false
}

// ConsumeChar will capture a single character from the options, returns the current char and a boolean (is the character one of the options)
func (scan *Scan) ConsumeChar(options ...byte) bool {
	_, found := scan.CaptureChar(options...)
	return found
}

// CaptureUntil will search the scan until it reaches a delimeter option
func (scan *Scan) CaptureUntil(options ...byte) ([]byte, byte, bool) {
	captured := []byte{}
	for scan.pos < len(scan.body) {
		currentChar := scan.body[scan.pos]
		scan.pos++
		for _, option := range options {
			if option == currentChar {
				return captured, currentChar, true
			}
		}
		captured = append(captured, currentChar)
	}
	return captured, '\n', false
}

//SetUriType gives the ability to set URI type e.g. sip:, sips:
func (sf *SipTo) SetUriType(uriType string) {
	sf.UriType = uriType
}

//SetUser set user portion of uri
func (sf *SipTo) SetUser(user string) {
	sf.User = []byte(user)
}

//SetHost set host protion of uri
func (sf *SipTo) SetHost(host string) {
	sf.Host = []byte(host)
}

//SetPort set port portion of uri
func (sf *SipTo) SetPort(port string) {
	sf.Port = []byte(port)
}

//SetUserType sets user type
func (sf *SipTo) SetUserType(userType string) {
	sf.UserType = []byte(userType)
}

//SetTag sets To Tag
func (sf *SipTo) SetTag(tag string) {
	sf.Tag = []byte(tag)
}

//String returns header as string
func (sf *SipTo) String() string {
	requestline := "To: "
	requestline += "<" + sf.UriType + ":" + string(sf.User) + "@" + string(sf.Host) + ">"
	if sf.Tag != nil {
		requestline += ";tag=" + string(sf.Tag)
	}
	return requestline
}

//ParseSipTo parses SIP To Header
func ParseSipTo(v []byte, out *SipTo) {

	pos := 0
	state := fieldBase

	// Init the output area
	out.UriType = ""
	out.Name = nil
	out.User = nil
	out.Host = nil
	out.Port = nil
	out.Tag = nil
	out.UserType = nil
	out.Src = nil

	// Keep the source line if needed
	if keepSrc {
		out.Src = v
	}

	scan := NewScan(v)

	// Loop through the bytes making up the line
	for pos < len(v) {
		// FSM
		switch state {
		case fieldBase:
			state = handleFieldBaseState(scan, out)

		case fieldNameQ:
			capturedString, _, _ := scan.CaptureUntil('"')
			state = fieldBase
			out.Name = capturedString
			continue

		case fieldName:
			capturedString, _, _ := scan.CaptureUntil('<', ' ')
			state = fieldBase
			out.Name = capturedString
			continue

		case fieldUser:
			capturedString, _, _ := scan.CaptureUntil('@')
			state = fieldUserHost
			out.User = capturedString
			continue

		case fieldUserHost:
			capturedString, last, _ := scan.CaptureUntil(':', ';', '>')
			out.Host = capturedString
			state = fieldBase
			if last == ':' {
				state = fieldPort
			}
			continue

		case fieldPort:
			capturedString, _, _ := scan.CaptureUntil(';', '>', ' ')
			state = fieldBase
			out.Port = capturedString
			continue

		case fieldUserType:
			capturedString, _, _ := scan.CaptureUntil(';', '>', ' ')
			state = fieldBase
			out.UserType = capturedString
			continue

		case fieldTag:
			capturedString, _, _ := scan.CaptureUntil(';', '>', ' ')
			state = fieldBase
			out.Tag = capturedString
			continue

		case fieldIgnore:
			scan.CaptureUntil(';', '>')
			state = fieldBase
			continue
		}
		pos++
	}
}

func handleFieldBaseState(scan *Scan, out *SipTo) int {
	if scan.ConsumeChar('"') && out.UriType == "" {
		return fieldNameQ
	}
	if scan.ConsumeChar(' ') {
		// Not a space so check for uri types
		protocol, hasProtocol := scan.CaptureString("sip:", "sips:", "tel:")
		if hasProtocol {
			out.UriType = protocol[:len(protocol)-1]
			return fieldUser
		}

		// Look for a Tag identifier
		if scan.ConsumeString("tag=") {
			return fieldTag
		}
		// Look for a User Type identifier
		if scan.ConsumeString("user=") {
			return fieldUserType
		}
		// Look for other identifiers and ignore
		if scan.ConsumeChar('=') {
			return fieldIgnore
		}
		// Check for other chrs
		if !scan.CheckChar('<', '>', ';') && out.UriType == "" {
			return fieldName
		}
	}
	return fieldBase

}
