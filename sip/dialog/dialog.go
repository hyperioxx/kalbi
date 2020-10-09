package dialog

import (
	"github.com/KalbiProject/Kalbi/sip/transaction"
)





func NewDialog() *Dialog {
	diag := new(Dialog)
	diag.DialogId = GenerateDialogId()
	return diag
}

type Dialog struct {
	DialogId int32
	CallId   string
	To_tag   string
	From_tag string
	ServerTx transaction.Transaction
	ClientTx transaction.Transaction
	Cseq     uint32
}
