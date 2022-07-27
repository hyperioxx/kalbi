package kalbi

import (
	"fmt"

	"github.com/KalbiProject/kalbi/log"
	"github.com/KalbiProject/kalbi/sip/dialog"
	"github.com/KalbiProject/kalbi/sip/message"
	"github.com/KalbiProject/kalbi/sip/transaction"
	"github.com/KalbiProject/kalbi/transport"
)

//NewSipStack  creates new sip stack
func NewSipStack(Name string) *SipStack {
	stack := new(SipStack)
	stack.Name = Name
	stack.TransManager = transaction.NewTransactionManager()
	stack.TransportChannel = make(chan message.SipEventObject)
	stack.funcMap = make(map[string]func(message.SipEventObject))

	return stack
}

//SipStack has multiple protocol listning points
type SipStack struct {
	Name             string
	ListeningPoints  []message.ListeningPoint
	OutputPoint      chan message.SipMsg
	InputPoint       chan message.SipMsg
	Alive            bool
	TransManager     *transaction.TransactionManager
	Dialogs          []dialog.Dialog
	TransportChannel chan message.SipEventObject
	sipListener      message.SipListener
	funcMap          map[string]func(message.SipEventObject)
}

//GetTransactionManager returns TransactionManager
func (ed *SipStack) GetTransactionManager() *transaction.TransactionManager {
	return ed.TransManager
}

//CreateListenPoint creates listening point to the event dispatcher
func (ed *SipStack) CreateListenPoint(protocol string, host string, port int) message.ListeningPoint {
	listenpoint := transport.NewTransportListenPoint(protocol, host, port)
	listenpoint.SetTransportChannel(ed.TransportChannel)
	ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
	return listenpoint
}

//SetSipListener sets a struct that follows the SipListener interface
func (ed *SipStack) SetSipListener(listener message.SipListener) {
	ed.sipListener = listener
}

//INVITE used to attach a callback function when an INVITE is received
func (ed *SipStack) INVITE(handler func(message.SipEventObject)) {
	ed.funcMap["INVITE"] = handler
}

//ACK used to attach a callback function when an ACK is received
func (ed *SipStack) ACK(handler func(message.SipEventObject)) {
	ed.funcMap["ACK"] = handler
}

//BYE used to attach a callback function when an BYE is received
func (ed *SipStack) BYE(handler func(message.SipEventObject)) {
	ed.funcMap["BYE"] = handler
}

//CANCEL used to attach a callback function when an CANCEL is received
func (ed *SipStack) CANCEL(handler func(message.SipEventObject)) {
	ed.funcMap["CANCEL"] = handler
}

//REGISTER used to attach a callback function when an REGISTER is received
func (ed *SipStack) REGISTER(handler func(message.SipEventObject)) {
	ed.funcMap["REGISTER"] = handler
}

//INFO used to attach a callback function when an INFO is received
func (ed *SipStack) INFO(handler func(message.SipEventObject)) {
	ed.funcMap["INFO"] = handler
}

//OPTIONS used to attach a callback function when an OPTIONS is received
func (ed *SipStack) OPTIONS(handler func(message.SipEventObject)) {
	ed.funcMap["OPTIONS"] = handler
}

//PRACK used to attach a callback function when an PRACK is received
func (ed *SipStack) PRACK(handler func(message.SipEventObject)) {
	ed.funcMap["PRACK"] = handler
}

//SUBSCRIBE used to attach a callback function when an SUBSCRIBE is received
func (ed *SipStack) SUBSCRIBE(handler func(message.SipEventObject)) {
	ed.funcMap["SUBSCRIBE"] = handler
}

//NOTIFY used to attach a callback function when an NOTIFY is received
func (ed *SipStack) NOTIFY(handler func(message.SipEventObject)) {
	ed.funcMap["NOTIFY"] = handler
}

//PUBLISH used to attach a callback function when an PUBLISH is received
func (ed *SipStack) PUBLISH(handler func(message.SipEventObject)) {
	ed.funcMap["PUBLISH"] = handler
}

//REFER used to attach a callback function when an REFER is received
func (ed *SipStack) REFER(handler func(message.SipEventObject)) {
	ed.funcMap["REFER"] = handler
}

//MESSAGE used to attach a callback function when an MESSAGE is received
func (ed *SipStack) MESSAGE(handler func(message.SipEventObject)) {
	ed.funcMap["MESSAGE"] = handler
}

//UPDATE used to attach a callback function when an UPDATE is received
func (ed *SipStack) UPDATE(handler func(message.SipEventObject)) {
	ed.funcMap["UPDATE"] = handler
}

//IsAlive check if SipStack is alive
func (ed *SipStack) IsAlive() bool {
	return ed.Alive
}

//Stop stops SipStack execution
func (ed *SipStack) Stop() {
	log.Log.Info("Stopping SIP Server...")
	ed.Alive = false
}

//Start starts the sip stack
func (ed *SipStack) Start() {
	log.Log.Info("Starting SIP Server...")
	ed.TransManager.ListeningPoint = ed.ListeningPoints[0]
	ed.Alive = true
	for _, listeningPoint := range ed.ListeningPoints {
		go listeningPoint.Start()
	}

	for ed.Alive {
		msg := <-ed.TransportChannel
		event := ed.TransManager.Handle(msg)
		message := event.GetSipMessage()
		fmt.Print(message)
		if message.Req.StatusCode != nil {
			go ed.sipListener.HandleResponses(event)
		} else if message.Req.Method != nil {
			switch string(message.Req.Method) {
			case "INVITE":
				ed.funcMap["INVITE"](event)
			case "ACK":
				ed.funcMap["ACK"](event)
			case "BYE":
				ed.funcMap["BYE"](event)
			case "CANCEL":
				ed.funcMap["CANCEL"](event)
			case "REGISTER":
				ed.funcMap["REGISTER"](event)

			}
		}
		//TODO: Handle failed SIP parse send 400 Bad Request

	}
}
