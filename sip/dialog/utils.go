package dialog

import (
	"math/rand"
	"time"
)

//GenerateDialogId creates new dialog ID
func GenerateDialogId() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}
