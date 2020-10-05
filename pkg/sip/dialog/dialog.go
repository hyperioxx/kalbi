package dialog

import "Kalbi/pkg/sip/transaction"


type Dialog struct {
	CallId    string
	To_tag    string 
	From_tag  string 
	CurrentTx transaction.Transaction
	Cseq      uint32
}
