package dispatcher

import "Kalbi/transport"
import "Kalbi/log"


//EventDispatcher has multiple protocol listning points 
type EventDispatcher struct {
    ListeningPoints []transport.ListeningPoint
}

//AddListenPoint adds listening point to the event dispatcher
func (ed *EventDispatcher) AddListenPoint(listenpoint transport.ListeningPoint){
    ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
}

//Start starts the event dispatccher
func (ed *EventDispatcher) Start(){
	log.Log.Info("Starting Kalbi server")

}