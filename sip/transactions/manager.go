package transaction

//TransactionManager handles SIP transactions
type TransactionManager struct {
	tx map[string] Transaction
}

func (tm *TransactionManager) init(){
	tm.tx = make(map[string] Transaction)
}

//SetChannels setup input & output channels for 
func (tm *TransactionManager) SetChannels(){
    
}

// Start runs TransManager
func (tm *TransactionManager) Start(){

}





