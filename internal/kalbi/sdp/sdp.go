package sdp

import (
	"fmt"
	"github.com/marv2097/siprocket"
)

func HandleSdp(Sdp siprocket.SdpMsg) error {
	fmt.Println(string(Sdp.ConnData.Src))

	return nil

}
