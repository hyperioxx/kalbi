package stack

//import "fmt"
import (
	"fmt"
	"Kalbi/pkg/sip/events"
	"Kalbi/pkg/sip/message"
	"Kalbi/pkg/transport"
	"Kalbi/pkg/sip/transaction"
	"Kalbi/pkg/log"
)

//NewSipStack  creates new sip stack
func NewSipStack(Name string) *SipStack {
	stack := new(SipStack)
	stack.Name = Name
	stack.TxMng = transaction.NewTransactionManager()
	return stack
}


//SipStack has multiple protocol listning points
type SipStack struct {
	Name              string
	ListeningPoints   []transport.ListeningPoint
	OutputPoint       chan message.SipMsg
	InputPoint        chan message.SipMsg
	Alive             bool
	TxMng             *transaction.TransactionManager
	EventChannelList  []chan events.Event

}


//CreateListenPoint creates listening point to the event dispatcher
func (ed *SipStack) CreateListenPoint(protocol string, host string, port int) transport.ListeningPoint {
	listenpoint := transport.NewTransportListenPoint(protocol, host, port)
	ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
    return listenpoint
}

func (ed *SipStack) CreateEventsChannel() chan events.Event {
	eventChannel := make(chan events.Event)
	ed.EventChannelList = append(ed.EventChannelList, eventChannel)
	return eventChannel 
}


func (ed *SipStack) IsAlive() bool {
    return ed.Alive
}


//Start starts the sip stack
func (ed *SipStack) Start() {
	log.Log.Info("Starting Sip Stack")
    ed.Alive = true
	for ed.Alive == true {
		for _, listeningPoint := range ed.ListeningPoints {
			fmt.Print(listeningPoint)
			msg := listeningPoint.Read()
			fmt.Println(msg)
		    ed.TxMng.Handle(msg)
			event := events.NewEvent(*msg, ed.TxMng.FindTransaction(string(msg.Via[0].Branch)))
			for _, eventchannel := range ed.EventChannelList{
				  eventchannel <- *event
			}
		}
       
	}
}
