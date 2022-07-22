package dialog

import (
	"math/rand"
	"time"
)

//GenerateDialogID creates new dialog ID
func GenerateDialogID() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}
