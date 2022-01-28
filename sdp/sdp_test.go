package sdp

import (
	"fmt"
	"reflect"
	"testing"
)

type fields struct {
	Origin    SdpOrigin
	Version   sdpVersion
	Time      sdpTime
	MediaDesc sdpMediaDesc
	Attrib    []sdpAttrib
	ConnData  sdpConnData
}

var (
	sdpAtts = []sdpAttrib{
		{Cat: []byte{116, 101, 115, 116},
			Val: []byte{116, 101, 115, 116},
			Src: []byte{116, 101, 115, 116}},
		{Cat: []byte{116, 101, 115, 116},
			Val: []byte{116, 101, 115, 116},
			Src: []byte{116, 101, 115, 116}}}

	testSdpOrigin = SdpOrigin{
		Username:       []byte{116, 101, 115, 116},
		SessionID:      []byte{116, 101, 115, 116},
		SessionVersion: []byte{116, 101, 115, 116},
		NetType:        []byte{116, 101, 115, 116},
		AddrType:       []byte{116, 101, 115, 116},
		UniAddr:        []byte{116, 101, 115, 116},
		Src:            []byte{116, 101, 115, 116}}

	testOne = fields{Origin: testSdpOrigin,
		Version:   sdpVersion{[]byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}},
		Time:      sdpTime{[]byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}},
		MediaDesc: sdpMediaDesc{[]byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}},
		Attrib:    sdpAtts,
		ConnData:  sdpConnData{[]byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}},
	}
	testOneBytes = []byte{118, 61, 116, 101, 115, 116, 13, 10, 111, 61, 116, 101, 115, 116, 32, 116, 101,
		115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115,
		116, 13, 10, 115, 61, 116, 101, 115, 116, 13, 10, 99, 61, 116, 101, 115, 116, 32, 116, 101, 115, 116,
		32, 116, 101, 115, 116, 13, 10, 116, 61, 116, 101, 115, 116, 32, 116, 101, 115, 116, 13, 10, 109, 61,
		116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 13, 10, 97,
		61, 116, 101, 115, 116, 58, 116, 101, 115, 116, 13, 10, 97, 61, 116, 101, 115, 116, 58, 116, 101, 115, 116, 13, 10}

	testInBytes = []byte{116, 101, 115, 116}

	testInBytesAndSpaces = []byte{116, 101, 115, 116, 32, 32, 32, 32, 32, 32}
	_                    = testInBytesAndSpaces

	testParseBytes = []byte{118, 61, 116, 101, 115, 116, 13, 10, 111, 61, 116, 101, 115, 116, 32, 116, 101,
		115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115,
		116, 13, 10, 115, 61, 116, 101, 115, 116, 13, 10, 99, 61, 116, 101, 115, 116, 32, 116, 101, 115, 116,
		32, 116, 101, 115, 116, 13, 10, 116, 61, 116, 101, 115, 116, 32, 116, 101, 115, 116, 13, 10, 109, 61,
		116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 13, 10, 97,
		61, 116, 101, 115, 116, 58, 116, 101, 115, 116, 13, 10, 97, 61, 116, 101, 115, 116, 58, 116, 101, 115, 116, 13, 10}

	testSpdMsgOrigin = SdpOrigin{
		Username:       []byte{116, 101, 115, 116},
		SessionID:      []byte{116, 101, 115, 116},
		SessionVersion: []byte{116, 101, 115, 116},
		NetType:        []byte{116, 101, 115, 116},
		AddrType:       []byte{116, 101, 115, 116},
		UniAddr:        []byte{116, 101, 115, 116},
		Src:            []byte{116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116}}

	testSpdMsgApdAtts = []sdpAttrib{
		{Cat: []byte{116, 101, 115, 116},
			Val: []byte{116, 101, 115, 116},
			Src: []byte{116, 101, 115, 116, 58, 116, 101, 115, 116}},
		{Cat: []byte{116, 101, 115, 116},
			Val: []byte{116, 101, 115, 116},
			Src: []byte{116, 101, 115, 116, 58, 116, 101, 115, 116}}}

	testSpdMsgExpected = SdpMsg{
		Origin:  testSpdMsgOrigin,
		Version: sdpVersion{[]byte{116, 101, 115, 116}, []byte{118, 61, 116, 101, 115, 116}},
		Time:    sdpTime{[]byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116, 32, 116, 101, 115, 116}},
		MediaDesc: sdpMediaDesc{[]byte{116, 101, 115, 116},
			[]byte{116, 101, 115, 116},
			[]byte{116, 101, 115, 116},
			[]byte{116, 101, 115, 116},
			[]byte{116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116}},
		Attrib:   testSpdMsgApdAtts,
		ConnData: sdpConnData{[]byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116}, []byte{116, 101, 115, 116, 32, 116, 101, 115, 116, 32, 116, 101, 115, 116}},
	}
)

func testByteEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSdpMsg_Size(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"testOne - 129", testOne, 129},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SdpMsg{
				Origin:    tt.fields.Origin,
				Version:   tt.fields.Version,
				Time:      tt.fields.Time,
				MediaDesc: tt.fields.MediaDesc,
				Attrib:    tt.fields.Attrib,
				ConnData:  tt.fields.ConnData,
			}
			if got := sm.Size(); got != tt.want {
				t.Errorf("SdpMsg.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSdpMsg_String(t *testing.T) {

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"testOne - 66", testOne, testOneBytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SdpMsg{
				Origin:    tt.fields.Origin,
				Version:   tt.fields.Version,
				Time:      tt.fields.Time,
				MediaDesc: tt.fields.MediaDesc,
				Attrib:    tt.fields.Attrib,
				ConnData:  tt.fields.ConnData,
			}

			got := []byte(sm.String())
			if !testByteEq(got, tt.want) {
				t.Errorf("SdpMsg.String() in []byte = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexSep(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 byte
	}{
		{"testOne - 32", args{testInBytes}, -1, byte(32)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := indexSep(tt.args.s)
			if got != tt.want {
				t.Errorf("indexSep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("indexSep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		v []byte
	}
	tests := []struct {
		name       string
		args       args
		wantOutput SdpMsg
	}{
		{"testOne - 32", args{testParseBytes}, testSpdMsgExpected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Parse(tt.args.v); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				fmt.Println("gotOutput.ConnData", gotOutput.Version)
				fmt.Println("gotOutput.ConnData", gotOutput.Version.Src)
				fmt.Println("gotOutput.ConnData", gotOutput.Version.Val)
				t.Errorf("Parse() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_testByteEq(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"testing test function", args{a: []byte{116, 101, 115, 116}, b: []byte{116, 101, 115, 116}}, true},
		{"testing test function", args{a: []byte{116, 101, 115, 116}, b: []byte{11, 101, 115, 116}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testByteEq(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("testByteEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
