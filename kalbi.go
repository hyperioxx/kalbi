package kalbi

import (
	"kalbi/sip/dialog"
	"kalbi/sip/message"
	"kalbi/sip/transaction"
	"strings"
)

// NewSipStack  creates new sip stack
func NewSipStack(Name string) *SipStack {
	stack := new(SipStack)
	stack.Name = Name
	stack.TransManager = transaction.NewTransactionManager()
	stack.TransportChannel = make(chan message.SipEventObject)
	stack.funcMap = make(map[string]func(message.SipEventObject))

	return stack
}

// SipStack has multiple protocol listning points
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

// SetSipListener sets a struct that follows the SipListener interface
func (ed *SipStack) SetSipListener(listener message.SipListener) {
	ed.sipListener = listener
}

func (s *SipStack) Handle(method string, handler func(message.SipEventObject)) {
	//TODO: check if supported method
	s.funcMap[method] = handler
}

// IsAlive check if SipStack is alive
func (ed *SipStack) IsAlive() bool {
	return ed.Alive
}

// Stop stops SipStack execution
func (ed *SipStack) Stop() {
	ed.Alive = false
}

// Start starts the sip stack
func (ed *SipStack) Start() {
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
			method := strings.ToUpper(string(message.Req.Method))
			handler := ed.funcMap[method]
			handler(event)
		}
	}
}
