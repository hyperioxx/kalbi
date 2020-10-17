package message

/*

20.7 Authorization

   The Authorization header field contains authentication credentials of
   a UA.  Section 22.2 overviews the use of the Authorization header
   field, and Section 22.4 describes the syntax and semantics when used
   with HTTP authentication.

   This header field, along with Proxy-Authorization, breaks the general
   rules about multiple header field values.  Although not a comma-
   separated list, this header field name may be present multiple times,
   and MUST NOT be combined into a single header line using the usual
   rules described in Section 7.3.

*/

//SipAuth SIP Authorization Header
type SipAuth struct {
	Username  []byte
	Realm     []byte
	Nonce     []byte
	CNonce    []byte
	QoP       []byte
	Algorithm []byte
	Nc        []byte
	URI       []byte
	Response  []byte
	Src       []byte
}

//ParseSipAuth parse's WWW-Authenticate/Authorization headers
func ParseSipAuth(v []byte, out *SipAuth) {

	pos := 0
	state := fieldBase

	// Init the output area
	out.Username = nil
	out.Realm = nil
	out.Nonce = nil
	out.CNonce = nil
	out.QoP = nil
	out.Algorithm = nil
	out.Nc = nil
	out.Response = nil
	out.Src = nil

	// Keep the source line if needed
	if keepSrc {
		out.Src = v
	}

	// Loop through the bytes making up the line
	for pos < len(v) {
		// FSM

		switch state {
		case fieldBase:
			if v[pos] != ',' || v[pos] != ' ' {

				if getString(v, pos, pos+9) == "username=" {
					state = fieldUser
					if v[pos+9] == '"' {
						pos = pos + 10
					} else {
						pos = pos + 9
					}
					continue
				}

				if getString(v, pos, pos+9) == "response=" {
					state = fieldResponse
					if v[pos+9] == '"' {
						pos = pos + 10
					} else {
						pos = pos + 9
					}
					continue

				}

				if getString(v, pos, pos+4) == "qop=" {
					state = fieldQop
					if v[pos+4] == '"' {
						pos = pos + 5
					} else {
						pos = pos + 4
					}
					continue
				}

				if getString(v, pos, pos+4) == "uri=" {
					state = fieldURI
					if v[pos+4] == '"' {
						pos = pos + 5
					} else {
						pos = pos + 4
					}
					continue
				}

				if getString(v, pos, pos+3) == "nc=" {
					state = fieldNC
					pos = pos + 3
					continue
				}

				if getString(v, pos, pos+7) == "cnonce=" {
					state = fieldCNonce
					pos = pos + 8
					continue
				}

				if getString(v, pos, pos+6) == "nonce=" {
					state = fieldNonce
					pos = pos + 7
					continue
				}
				if getString(v, pos, pos+6) == "realm=" {
					state = fieldRealm
					pos = pos + 7
					continue
				}
				if getString(v, pos, pos+10) == "algorithm=" {
					state = fieldAlgorithm
					pos = pos + 10
					continue
				}

			}

		case fieldQop:
			if v[pos] == ' ' || v[pos] == ',' || v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.QoP = append(out.QoP, v[pos])

		case fieldNonce:
			if v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.Nonce = append(out.Nonce, v[pos])

		case fieldCNonce:
			if v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.CNonce = append(out.CNonce, v[pos])

		case fieldURI:
			if v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.URI = append(out.URI, v[pos])

		case fieldResponse:
			if v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.Response = append(out.Response, v[pos])

		case fieldRealm:
			if v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.Realm = append(out.Realm, v[pos])

		case fieldAlgorithm:
			if v[pos] == ' ' || v[pos] == ',' {
				state = fieldBase
				pos++
				continue
			}
			out.Algorithm = append(out.Algorithm, v[pos])

		case fieldUser:
			if v[pos] == '"' {
				state = fieldBase
				pos++
				continue
			}
			out.Username = append(out.Username, v[pos])

		case fieldNC:
			if v[pos] == ',' || v[pos] == ' ' {
				state = fieldBase
				pos++
				continue
			}
			out.Nc = append(out.Nc, v[pos])

		case fieldIgnore:
			if v[pos] == ' ' || v[pos] == ',' {
				state = fieldBase
				pos++
				continue
			}

		}
		pos++
	}
}
