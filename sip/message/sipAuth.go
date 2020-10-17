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
	state := FIELD_BASE

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
		case FIELD_BASE:
			if v[pos] != ',' || v[pos] != ' ' {

				if getString(v, pos, pos+9) == "username=" {
					state = FIELD_USER
					if v[pos+9] == '"' {
						pos = pos + 10
					} else {
						pos = pos + 9
					}
					continue
				}

				if getString(v, pos, pos+9) == "response=" {
					state = FIELD_RESPONSE
					if v[pos+9] == '"' {
						pos = pos + 10
					} else {
						pos = pos + 9
					}
					continue

				}

				if getString(v, pos, pos+4) == "qop=" {
					state = FIELD_QOP
					if v[pos+4] == '"' {
						pos = pos + 5
					} else {
						pos = pos + 4
					}
					continue
				}

				if getString(v, pos, pos+4) == "uri=" {
					state = FIELD_URI
					if v[pos+4] == '"' {
						pos = pos + 5
					} else {
						pos = pos + 4
					}
					continue
				}

				if getString(v, pos, pos+3) == "nc=" {
					state = FIELD_NC
					pos = pos + 3
					continue
				}

				if getString(v, pos, pos+7) == "cnonce=" {
					state = FIELD_CNONCE
					pos = pos + 8
					continue
				}

				if getString(v, pos, pos+6) == "nonce=" {
					state = FIELD_NONCE
					pos = pos + 7
					continue
				}
				if getString(v, pos, pos+6) == "realm=" {
					state = FIELD_REALM
					pos = pos + 7
					continue
				}
				if getString(v, pos, pos+10) == "algorithm=" {
					state = FIELD_ALGORITHM
					pos = pos + 10
					continue
				}

			}

		case FIELD_QOP:
			if v[pos] == ' ' || v[pos] == ',' || v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.QoP = append(out.QoP, v[pos])

		case FIELD_NONCE:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Nonce = append(out.Nonce, v[pos])

		case FIELD_CNONCE:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.CNonce = append(out.CNonce, v[pos])

		case FIELD_URI:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.URI = append(out.URI, v[pos])

		case FIELD_RESPONSE:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Response = append(out.Response, v[pos])

		case FIELD_REALM:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Realm = append(out.Realm, v[pos])

		case FIELD_ALGORITHM:
			if v[pos] == ' ' || v[pos] == ',' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Algorithm = append(out.Algorithm, v[pos])

		case FIELD_USER:
			if v[pos] == '"' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Username = append(out.Username, v[pos])

		case FIELD_NC:
			if v[pos] == ',' || v[pos] == ' ' {
				state = FIELD_BASE
				pos++
				continue
			}
			out.Nc = append(out.Nc, v[pos])

		case FIELD_IGNORE:
			if v[pos] == ' ' || v[pos] == ',' {
				state = FIELD_BASE
				pos++
				continue
			}

		}
		pos++
	}
}
