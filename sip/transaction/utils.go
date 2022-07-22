package transaction

import (
	"fmt"
	"math/rand"

	"github.com/KalbiProject/kalbi/log"
)

//GenerateBranchID creates new branch ID
func GenerateBranchID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x", b[0:4], b[4:6])
	return "z9hG4bK-" + uuid
}
