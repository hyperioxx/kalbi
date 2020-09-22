package dispatcher

//import "fmt"
import "github.com/marv2097/siprocket"
import "Kalbi/transport"
import "Kalbi/log"



//EventDispatcher has multiple protocol listning points 
type EventDispatcher struct {
    ListeningPoints []transport.ListeningPoint
    OutputPoints [] chan siprocket.SipMsg
}

//AddListenPoint adds listening point to the event dispatcher
func (ed *EventDispatcher) AddListenPoint(listenpoint transport.ListeningPoint){
    ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
}

//AddChannel give the ability to add channels for callbacks on each request
func (ed *EventDispatcher) AddChannel(c chan siprocket.SipMsg){
     ed.OutputPoints = append(ed.OutputPoints, c)
}

//Start starts the event dispatccher
func (ed *EventDispatcher) Start(){
    log.Log.Info("Starting Kalbi server")

    for {
        for _, listeningPoint := range ed.ListeningPoints {
            msg := listeningPoint.Read()
            for _ , OutputPoint := range ed.OutputPoints{
                OutputPoint <- *msg
            } 
            
        }
    
    }


}