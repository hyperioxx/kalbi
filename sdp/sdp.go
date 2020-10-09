package sdp

import ("bytes"
		"strings")


var keep_src = true

const FIELD_NULL = 0
const FIELD_BASE = 1
const FIELD_VALUE = 2
const FIELD_NAME = 3
const FIELD_NAMEQ = 4
const FIELD_USER = 5
const FIELD_HOST = 6
const FIELD_PORT = 7
const FIELD_TAG = 8
const FIELD_ID = 9
const FIELD_METHOD = 10
const FIELD_TRAN = 11
const FIELD_BRANCH = 12
const FIELD_RPORT = 13
const FIELD_MADDR = 14
const FIELD_TTL = 15
const FIELD_REC = 16
const FIELD_EXPIRES = 17
const FIELD_Q = 18
const FIELD_USERTYPE = 19
const FIELD_STATUS = 20
const FIELD_STATUSDESC = 21

const FIELD_ADDRTYPE = 40
const FIELD_CONNADDR = 41
const FIELD_MEDIA = 42
const FIELD_PROTO = 43
const FIELD_FMT = 44
const FIELD_CAT = 45

const FIELD_IGNORE = 255


//SdpMsg is struc
type SdpMsg struct {
	MediaDesc sdpMediaDesc
	Attrib    []sdpAttrib
	ConnData  sdpConnData
}
		
		

func indexSep(s []byte) (int, byte) {
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			return i, '='
		}
	}
	return -1, ' '
}

//Parse Parses SDP 
func Parse(v []byte) (output SdpMsg) {
	attr_idx := 0
	output.Attrib = make([]sdpAttrib, 0, 8)

	lines := bytes.Split(v, []byte("\r\n"))

	for _, line := range lines {
		spos, stype := indexSep(line)
		//fmt.Println(i, string(line))
		line = bytes.TrimSpace(line)

			if spos == 1 && stype == '=' {
				// SDP: Break up into header and value
				lhdr := strings.ToLower(string(line[0]))
				lval := bytes.TrimSpace(line[2:])
				// Switch on the line header
				//fmt.Println(i, spos, string(lhdr), string(lval))
				switch {
				case lhdr == "v":
					
				case lhdr == "m":
					parseSdpMediaDesc(lval, &output.MediaDesc)
				case lhdr == "c":
					parseSdpConnectionData(lval, &output.ConnData)
				case lhdr == "a":
					var tmpAttrib sdpAttrib
					output.Attrib = append(output.Attrib, tmpAttrib)
					parseSdpAttrib(lval, &output.Attrib[attr_idx])
					attr_idx++

				} // End of Switch

			}
		}
		return
	}

	