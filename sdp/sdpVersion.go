package sdp


type sdpVersion struct {

	Val []byte // Version number
	Src []byte // Full source if needed
}

func (sv *sdpVersion) Export() string {
	line := "v="
	line += string(sv.Val)
    return line
}
