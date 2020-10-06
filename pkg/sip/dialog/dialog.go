package dialog

import ("Kalbi/pkg/sip/transaction"
        "Kalbi/pkg/sip/utils")

func NewDialog() *Dialog{
	diag := new(Dialog)
	diag.DialogId = utils.GenerateDialogId()
	return diag
}


type Dialog struct {
	DialogId  int32
	CallId    string
	To_tag    string 
	From_tag  string 
	CurrentTx transaction.Transaction
	Cseq      uint32
}
