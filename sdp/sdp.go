package sdp

import (
	"bytes"

	"strings"
)

var keep_src = true

//FMS States
const FIELD_NULL = 0
const FIELD_BASE = 1
const FIELD_VALUE = 2
const FIELD_PORT = 3
const FIELD_ADDRTYPE = 4
const FIELD_CONNADDR = 5
const FIELD_MEDIA = 6
const FIELD_PROTO = 7
const FIELD_FMT = 8
const FIELD_CAT = 9
const FIELD_USERNAME = 10
const FIELD_SESSIONID = 11
const FIELD_SESSIONVERSION = 12
const FIELD_NETTYPE = 13
const FIELD_UNIADDR = 14
const FIELD_TIMESTART = 15
const FIELD_TIMESTOP = 16

const FIELD_IGNORE = 255

//SdpMsg is representation of an SDP message
type SdpMsg struct {
	Origin    SdpOrigin
	Version   sdpVersion
	Time      sdpTime
	MediaDesc sdpMediaDesc
	Attrib    []sdpAttrib
	ConnData  sdpConnData
}

//Size returns size in bytes
func (sm *SdpMsg) Size() int {
	sdp := sm.Export()
	return len([]byte(sdp))
}

func (sm *SdpMsg) Export() string {
	sdp := "\r\n"
	sdp += sm.Version.Export() + "\r\n"
	sdp += sm.Origin.Export() + "\r\n"
	sdp += sm.ConnData.Export() + "\r\n"
	sdp += sm.Time.Export() + "\r\n"
	for _, a := range sm.Attrib {
		sdp += a.Export() + "\r\n"
	}
	return sdp
	sdp += sm.MediaDesc.Export() + "\r\n"
	

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
				output.Version = sdpVersion{Val: lval, Src: line}
			case lhdr == "o":
				ParseSdpOrigin(lval, &output.Origin)
			case lhdr == "t":
				ParserSdpTime(lval, &output.Time)
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
