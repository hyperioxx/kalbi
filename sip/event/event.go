package event

import (
	"fmt"

	"github.com/KalbiProject/kalbi/interfaces"
	logger "github.com/KalbiProject/kalbi/logger"
	"github.com/KalbiProject/kalbi/sip/message"
)

//SipEvent object that gets passed to the SipListener
type SipEvent struct {
	sipmsg *message.SipMsg
	tx     interfaces.Transaction
	lp     interfaces.ListeningPoint
}

//GetSipMessage returns message that created this event
func (se *SipEvent) GetSipMessage() *message.SipMsg {
	return se.sipmsg
}

//SetSipMessage sets message that created this event
func (se *SipEvent) SetSipMessage(msg *message.SipMsg) {
	se.sipmsg = msg
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Method - %v", string(se.sipmsg.Req.Method)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.UriType - %v", string(se.sipmsg.Req.UriType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.StatusCode - %v", string(se.sipmsg.Req.StatusCode)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.StatusDesc - %v", string(se.sipmsg.Req.StatusDesc)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.User - %v", string(se.sipmsg.Req.User)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Host - %v", string(se.sipmsg.Req.Host)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Port - %v", string(se.sipmsg.Req.Port)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.UserType - %v", string(se.sipmsg.Req.UserType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Req.Src - %v", string(se.sipmsg.Req.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.UriType - %v", se.sipmsg.From.UriType))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Name - %v", string(se.sipmsg.From.Name)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.User - %v", string(se.sipmsg.From.User)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Host - %v", string(se.sipmsg.From.Host)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Port - %v", string(se.sipmsg.From.Port)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Tag - %v", string(se.sipmsg.From.Tag)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.UserType - %v", string(se.sipmsg.From.UserType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.From.Src - %v", string(se.sipmsg.From.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Host - %v", string(se.sipmsg.To.Host)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Name - %v", string(se.sipmsg.To.Name)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Port - %v", string(se.sipmsg.To.Port)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Src - %v", string(se.sipmsg.To.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.Tag - %v", string(se.sipmsg.To.Tag)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.UriType - %v", string(se.sipmsg.To.UriType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.User - %v", string(se.sipmsg.To.User)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.To.UserType - %v", string(se.sipmsg.To.UserType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Expires - %v", string(se.sipmsg.Contact.Expires)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Host - %v", string(se.sipmsg.Contact.Host)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Name - %v", string(se.sipmsg.Contact.Name)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Port - %v", string(se.sipmsg.Contact.Port)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Qval - %v", string(se.sipmsg.Contact.Qval)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Src - %v", string(se.sipmsg.Contact.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.Tran - %v", string(se.sipmsg.Contact.Tran)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.UriType - %v", string(se.sipmsg.Contact.UriType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Contact.User - %v", string(se.sipmsg.Contact.User)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Via - %v", se.sipmsg.Via))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Cseq.ID - %v", string(se.sipmsg.Cseq.ID)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Cseq.Method - %v", string(se.sipmsg.Cseq.Method)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Cseq.Src - %v", string(se.sipmsg.Cseq.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Algorithm - %v", string(se.sipmsg.Auth.Algorithm)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.CNonce - %v", string(se.sipmsg.Auth.CNonce)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Nc - %v", string(se.sipmsg.Auth.Nc)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Nonce - %v", string(se.sipmsg.Auth.Nonce)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.QoP - %v", string(se.sipmsg.Auth.QoP)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Realm - %v", string(se.sipmsg.Auth.Realm)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Response - %v", string(se.sipmsg.Auth.Response)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Src - %v", string(se.sipmsg.Auth.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.URI - %v", string(se.sipmsg.Auth.URI)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Auth.Username - %v", string(se.sipmsg.Auth.Username)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Ua.Src - %v", string(se.sipmsg.Ua.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Ua.Value - %v", string(se.sipmsg.Ua.Value)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Exp.Src - %v", string(se.sipmsg.Exp.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Exp.Value - %v", string(se.sipmsg.Exp.Value)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.MaxFwd.Src - %v", string(se.sipmsg.MaxFwd.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.MaxFwd.Value - %v", string(se.sipmsg.MaxFwd.Value)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.CallID.Src - %v", string(se.sipmsg.CallID.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.CallID.Value - %v", string(se.sipmsg.CallID.Value)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.ContType.Src - %v", string(se.sipmsg.ContType.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.ContType.Value - %v", string(se.sipmsg.ContType.Value)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.ContLen.Src - %v", string(se.sipmsg.ContLen.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.ContLen.Value - %v", string(se.sipmsg.ContLen.Value)))

	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Src - %v", string(se.sipmsg.Src)))

	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Body - %v", se.sipmsg.Body))

	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Attrib - %v", se.sipmsg.Sdp.Attrib))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.ConnData - %v", se.sipmsg.Sdp.ConnData))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.MediaDesc - %v", se.sipmsg.Sdp.MediaDesc))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.AddrType - %v", string(se.sipmsg.Sdp.Origin.AddrType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.NetType - %v", string(se.sipmsg.Sdp.Origin.NetType)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.SessionID - %v", string(se.sipmsg.Sdp.Origin.SessionID)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.SessionVersion - %v", string(se.sipmsg.Sdp.Origin.SessionVersion)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.Src - %v", string(se.sipmsg.Sdp.Origin.Src)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.UniAddr - %v", string(se.sipmsg.Sdp.Origin.UniAddr)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Origin.Username - %v", string(se.sipmsg.Sdp.Origin.Username)))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Time - %v", se.sipmsg.Sdp.Time))
	logger.Debug(fmt.Sprintf("SetSipMessage() se.sipmsg.Sdp.Version - %v", se.sipmsg.Sdp.Version))

	logger.Debug(fmt.Sprintf("SetSipMessage() msg - %v", msg))
}

//GetTransaction returns transaction related to the SIP message that created this event
func (se *SipEvent) GetTransaction() interfaces.Transaction {
	return se.tx
}

//SetTransaction sets transaction related to the SIP message that created this event
func (se *SipEvent) SetTransaction(tx interfaces.Transaction) {
	se.tx = tx
}

//SetListeningPoint gives ability to set interfaces.ListeningPoint
func (se *SipEvent) SetListeningPoint(lp interfaces.ListeningPoint) {
	se.lp = lp
}

//GetListeningPoint returns interfaces.ListeningPoint
func (se *SipEvent) GetListeningPoint() interfaces.ListeningPoint {
	return se.lp
}
