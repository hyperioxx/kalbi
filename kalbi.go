package kalbi

import (
	"github.com/KalbiProject/kalbi/interfaces"
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
	stack.TransportChannel = make(chan interfaces.SipEventObject)
	stack.funcMap = make(map[string]func(interfaces.SipEventObject))

	return stack
}

//SipStack has multiple protocol listning points
type SipStack struct {
	Name             string
	ListeningPoints  []interfaces.ListeningPoint
	OutputPoint      chan message.SipMsg
	InputPoint       chan message.SipMsg
	Alive            bool
	TransManager     *transaction.TransactionManager
	Dialogs          []dialog.Dialog
	TransportChannel chan interfaces.SipEventObject
	sipListener      interfaces.SipListener
	funcMap          map[string]func(interfaces.SipEventObject)
}

//GetTransactionManager returns TransactionManager
func (ed *SipStack) GetTransactionManager() *transaction.TransactionManager {
	return ed.TransManager
}

//CreateListenPoint creates listening point to the event dispatcher
func (ed *SipStack) CreateListenPoint(protocol string, host string, port int) interfaces.ListeningPoint {
	listenpoint := transport.NewTransportListenPoint(protocol, host, port)
	listenpoint.SetTransportChannel(ed.TransportChannel)
	ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
	return listenpoint
}

//SetSipListener sets a struct that follows the SipListener interface
func (ed *SipStack) SetSipListener(listener interfaces.SipListener) {
	ed.sipListener = listener
}

func (ed *SipStack) INVITE(handler func(interfaces.SipEventObject)) {
	ed.funcMap["INVITE"] = handler
}

func (ed *SipStack) ACK(handler func(interfaces.SipEventObject)) {
	ed.funcMap["ACK"] = handler
}

func (ed *SipStack) BYE(handler func(interfaces.SipEventObject)) {
	ed.funcMap["BYE"] = handler
}

func (ed *SipStack) CANCEL(handler func(interfaces.SipEventObject)) {
	ed.funcMap["CANCEL"] = handler
}

func (ed *SipStack) REGISTER(handler func(interfaces.SipEventObject)) {
	ed.funcMap["REGISTER"] = handler
}

func (ed *SipStack) INFO(handler func(interfaces.SipEventObject)) {
	ed.funcMap["INFO"] = handler
}

func (ed *SipStack) OPTIONS(handler func(interfaces.SipEventObject)) {
	ed.funcMap["OPTIONS"] = handler
}

func (ed *SipStack) PRACK(handler func(interfaces.SipEventObject)) {
	ed.funcMap["PRACK"] = handler
}

func (ed *SipStack) SUBSCRIBE(handler func(interfaces.SipEventObject)) {
	ed.funcMap["SUBSCRIBE"] = handler
}

func (ed *SipStack) NOTIFY(handler func(interfaces.SipEventObject)) {
	ed.funcMap["NOTIFY"] = handler
}

func (ed *SipStack) PUBLISH(handler func(interfaces.SipEventObject)) {
	ed.funcMap["PUBLISH"] = handler
}

func (ed *SipStack) REFER(handler func(interfaces.SipEventObject)) {
	ed.funcMap["REFER"] = handler
}

func (ed *SipStack) MESSAGE(handler func(interfaces.SipEventObject)) {
	ed.funcMap["MESSAGE"] = handler
}

func (ed *SipStack) UPDATE(handler func(interfaces.SipEventObject)) {
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
