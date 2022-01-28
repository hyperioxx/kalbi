# Kalbi

```go
import "github.com/KalbiProject/Kalbi"
```

## Index

- Kalbi
- [type SipStack](#type-sipstack)

  - [func NewSipStack(Name string) \*SipStack](#func-newsipstack)
  - [func (ed \*SipStack) CreateListenPoint(protocol string, host string, port int) transport.ListeningPoint](#func-sipstack-createlistenpoint)
  - [func (ed \*SipStack) CreateRequestsChannel() chan transaction.Transaction](#func-sipstack-createrequestschannel)
  - [func (ed \*SipStack) CreateResponseChannel() chan transaction.Transaction](#func-sipstack-createresponsechannel)
  - [func (ed *SipStack) GetTransactionManager() *transaction.TransactionManager](#func-sipstack-gettransactionmanager)
  - [func (ed \*SipStack) IsAlive() bool](#func-sipstack-isalive)
  - [func (ed \*SipStack) Start()](#func-sipstack-start)
  - [func (ed \*SipStack) Stop()](#func-sipstack-stop)

- sip

  - [func Parse(msg []byte) message.SipMsg](#func-parse)

- message

  - [Constants](#constants)
  - [func GenerateBranchId() string](#func-generatebranchid)
  - [func MessageDetails(data \*SipMsg) string](#func-messagedetails)
  - [func ParseSipContact(v []byte, out \*SipContact)](#func-parsesipcontact)
  - [func ParseSipCseq(v []byte, out \*SipCseq)](#func-parsesipcseq)
  - [func ParseSipFrom(v []byte, out \*SipFrom)](#func-parsesipfrom)
  - [func ParseSipReq(v []byte, out \*SipReq)](#func-parsesipreq)
  - [func ParseSipTo(v []byte, out \*SipTo)](#func-parsesipto)
  - [func ParseSipVia(v []byte, out \*SipVia)](#func-parsesipvia)
  - [type SipContact](#type-sipcontact)
    - [func (sc \*SipContact) String() string](#func-sipcontact-export)
    - [func (sc \*SipContact) SetName(name string)](#func-sipcontact-setname)
  - [type SipCseq](#type-sipcseq)
    - [func (sc \*SipCseq) String() string](#func-sipcseq-export)
    - [func (sc \*SipCseq) SetID(id string)](#func-sipcseq-setid)
    - [func (sc \*SipCseq) SetMethod(method string)](#func-sipcseq-setmethod)
  - [type SipFrom](#type-sipfrom)
    - [func (sf \*SipFrom) String() string](#func-sipfrom-export)
    - [func (sf \*SipFrom) SetHost(host string)](#func-sipfrom-sethost)
    - [func (sf \*SipFrom) SetPort(port string)](#func-sipfrom-setport)
    - [func (sf \*SipFrom) SetTag(tag string)](#func-sipfrom-settag)
    - [func (sf \*SipFrom) SetUriType(uriType string)](#func-sipfrom-seturitype)
    - [func (sf \*SipFrom) SetUser(user string)](#func-sipfrom-setuser)
    - [func (sf \*SipFrom) SetUserType(userType string)](#func-sipfrom-setusertype)
  - [type SipMsg](#type-sipmsg)
    - [func NewRequest(request string, to string, from string) \*SipMsg ](#func-newrequest)
    - [func NewResponse(response int, to string, from string) \*SipMsg ](#func-newresponse)
    - [func Parse(v []byte) (output SipMsg)](#func-parse)
    - [func (sm *SipMsg) CopyHeaders(msg *SipMsg)](#func-sipmsg-copyheaders)
    - [func (sm *SipMsg) CopySdp(msg *SipMsg)](#func-sipmsg-copysdp)
    - [func (sm \*SipMsg) String() string](#func-sipmsg-export)
    - [func (sm \*SipMsg) GetStatusCode() int](#func-sipmsg-getstatuscode)
  - [type SipReq](#type-sipreq)
    - [func (sr \*SipReq) String() string](#func-sipreq-export)
    - [func (sr \*SipReq) SetHost(host string)](#func-sipreq-sethost)
    - [func (sr \*SipReq) SetMethod(method string)](#func-sipreq-setmethod)
    - [func (sr \*SipReq) SetPort(port string)](#func-sipreq-setport)
    - [func (sr \*SipReq) SetStatusCode(code int)](#func-sipreq-setstatuscode)
    - [func (sr \*SipReq) SetStatusDesc(desc string)](#func-sipreq-setstatusdesc)
    - [func (sr \*SipReq) SetUriType(uriType string)](#func-sipreq-seturitype)
    - [func (sr \*SipReq) SetUser(user string)](#func-sipreq-setuser)
    - [func (sr \*SipReq) SetUserType(userType string)](#func-sipreq-setusertype)
  - [type SipTo](#type-sipto)
    - [func (sf \*SipTo) String() string](#func-sipto-export)
    - [func (sf \*SipTo) SetHost(host string)](#func-sipto-sethost)
    - [func (sf \*SipTo) SetPort(port string)](#func-sipto-setport)
    - [func (sf \*SipTo) SetTag(tag string)](#func-sipto-settag)
    - [func (sf \*SipTo) SetUriType(uriType string)](#func-sipto-seturitype)
    - [func (sf \*SipTo) SetUser(user string)](#func-sipto-setuser)
    - [func (sf \*SipTo) SetUserType(userType string)](#func-sipto-setusertype)
  - [type SipVal](#type-sipval)
    - [func (sv \*SipVal) String() string](#func-sipval-export)
    - [func (sv \*SipVal) SetValue(value string)](#func-sipval-setvalue)
  - [type SipVia](#type-sipvia)
    - [func (sv \*SipVia) String() string](#func-sipvia-export)
    - [func (sv \*SipVia) SetBranch(value string)](#func-sipvia-setbranch)
    - [func (sv \*SipVia) SetHost(value string)](#func-sipvia-sethost)
    - [func (sv \*SipVia) SetPort(value string)](#func-sipvia-setport)

- method
  - [Constants](#constants)
- status
  - [Constants](#constants)
  - [func StatusText(code int) string](#func-statustext)
- transaction

  - [Constants](#constants)
  - [func GenerateBranchId() string](#func-generatebranchid)
  - [type ClientTransaction](#type-clienttransaction)
    - [func (ct \*ClientTransaction) GetBranchID() string](#func-clienttransaction-getbranchid)
    - [func (ct *ClientTransaction) GetLastMessage() *message.SipMsg](#func-clienttransaction-getlastmessage)
    - [func (ct *ClientTransaction) GetOrigin() *message.SipMsg](#func-clienttransaction-getorigin)
    - [func (ct \*ClientTransaction) GetServerTransactionID() string](#func-clienttransaction-getservertransactionid)
    - [func (ct *ClientTransaction) InitFSM(msg *message.SipMsg)](#func-clienttransaction-initfsm)
    - [func (ct *ClientTransaction) Receive(msg *message.SipMsg)](#func-clienttransaction-receive)
    - [func (ct \*ClientTransaction) Resend()](#func-clienttransaction-resend)
    - [func (ct *ClientTransaction) Send(msg *message.SipMsg, host string, port string)](#func-clienttransaction-send)
    - [func (ct *ClientTransaction) SetLastMessage(msg *message.SipMsg)](#func-clienttransaction-setlastmessage)
    - [func (ct \*ClientTransaction) SetListeningPoint(lp transport.ListeningPoint)](#func-clienttransaction-setlisteningpoint)
    - [func (ct \*ClientTransaction) SetServerTransaction(txID string)](#func-clienttransaction-setservertransaction)
    - [func (ct *ClientTransaction) StatelessSend(msg *message.SipMsg, host string, port string)](#func-clienttransaction-statelesssend)
  - [type ServerTransaction](#type-servertransaction)
    - [func (st \*ServerTransaction) GetBranchID() string](#func-servertransaction-getbranchid)
    - [func (st *ServerTransaction) GetLastMessage() *message.SipMsg](#func-servertransaction-getlastmessage)
    - [func (st *ServerTransaction) GetOrigin() *message.SipMsg](#func-servertransaction-getorigin)
    - [func (st \*ServerTransaction) GetServerTransactionID() string](#func-servertransaction-getservertransactionid)
    - [func (st *ServerTransaction) InitFSM(msg *message.SipMsg)](#func-servertransaction-initfsm)
    - [func (st *ServerTransaction) Receive(msg *message.SipMsg)](#func-servertransaction-receive)
    - [func (st *ServerTransaction) Respond(msg *message.SipMsg)](#func-servertransaction-respond)
    - [func (st *ServerTransaction) Send(msg *message.SipMsg, host string, port string)](#func-servertransaction-send)
    - [func (st *ServerTransaction) SetLastMessage(msg *message.SipMsg)](#func-servertransaction-setlastmessage)
    - [func (st \*ServerTransaction) SetListeningPoint(lp transport.ListeningPoint)](#func-servertransaction-setlisteningpoint)
  - [type Transaction](#type-transaction)
  - [type TransactionManager](#type-transactionmanager)
    - [func NewTransactionManager() \*TransactionManager](#func-newtransactionmanager)
    - [func (tm \*TransactionManager) DeleteClientTransaction(branch string)](#func-transactionmanager-deleteclienttransaction)
    - [func (tm \*TransactionManager) DeleteServerTransaction(branch string)](#func-transactionmanager-deleteservertransaction)
    - [func (tm *TransactionManager) FindClientTransaction(msg *message.SipMsg) (Transaction, bool)](#func-transactionmanager-findclienttransaction)
    - [func (tm \*TransactionManager) FindClientTransactionByID(value string) (Transaction, bool)](#func-transactionmanager-findclienttransactionbyid)
    - [func (tm *TransactionManager) FindServerTransaction(msg *message.SipMsg) (Transaction, bool)](#func-transactionmanager-findservertransaction)
    - [func (tm \*TransactionManager) FindServerTransactionByID(value string) (Transaction, bool)](#func-transactionmanager-findservertransactionbyid)
    - [func (tm *TransactionManager) Handle(message *message.SipMsg)](#func-transactionmanager-handle)
    - [func (tm \*TransactionManager) MakeKey(msg message.SipMsg) string](#func-transactionmanager-makekey)
    - [func (tm *TransactionManager) NewClientTransaction(msg *message.SipMsg) \*ClientTransaction](#func-transactionmanager-newclienttransaction)
    - [func (tm *TransactionManager) NewServerTransaction(msg *message.SipMsg) \*ServerTransaction](#func-transactionmanager-newservertransaction)
    - [func (tm \*TransactionManager) PutClientTransaction(tx Transaction)](#func-transactionmanager-putclienttransaction)
    - [func (tm \*TransactionManager) PutServerTransaction(tx Transaction)](#func-transactionmanager-putservertransaction)

- sdp

  - [Constants](#constants)
  - [func ParseSdpOrigin(v []byte, out \*SdpOrigin)](#func-parsesdporigin)
  - [func ParserSdpTime(v []byte, out \*sdpTime)](#func-parsersdptime)
  - [type SdpMsg](#type-sdpmsg)
    - [func Parse(v []byte) (output SdpMsg)](#func-parse)
    - [func (sm \*SdpMsg) String() string](#func-sdpmsg-export)
    - [func (sm \*SdpMsg) Size() int](#func-sdpmsg-size)
  - [type SdpOrigin](#type-sdporigin)
    - [func (so \*SdpOrigin) String() string](#func-sdporigin-export)

- transport
  - [type ListeningPoint](#type-listeningpoint)
    - [func NewTransportListenPoint(protocol string, host string, port int) ListeningPoint](#func-newtransportlistenpoint)
  - [type TCPTransport](#type-tcptransport)
    - [func (tt \*TCPTransport) Build(host string, port int)](#func-tcptransport-build)
    - [func (tt *TCPTransport) Read() *message.SipMsg](#func-tcptransport-read)
  - [type UDPTransport](#type-udptransport)
    - [func (ut \*UDPTransport) Build(host string, port int)](#func-udptransport-build)
    - [func (ut *UDPTransport) Read() *message.SipMsg](#func-udptransport-read)
    - [func (ut \*UDPTransport) Send(host string, port string, msg string) error](#func-udptransport-send)

## type [SipStack](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L21-L31)

SipStack has multiple protocol listning points

```go
type SipStack struct {
    Name            string
    ListeningPoints []transport.ListeningPoint
    OutputPoint     chan message.SipMsg
    InputPoint      chan message.SipMsg
    Alive           bool
    TransManager    *transaction.TransactionManager
    Dialogs         []dialog.Dialog
    RequestChannel  chan transaction.Transaction
    ResponseChannel chan transaction.Transaction
}
```

### func [NewSipStack](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L12)

```go
func NewSipStack(Name string) *SipStack
```

NewSipStack creates new sip stack

### func (\*SipStack) [CreateListenPoint](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L38)

```go
func (ed *SipStack) CreateListenPoint(protocol string, host string, port int) transport.ListeningPoint
```

CreateListenPoint creates listening point to the event dispatcher

### func (\*SipStack) [CreateRequestsChannel](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L44)

```go
func (ed *SipStack) CreateRequestsChannel() chan transaction.Transaction
```

### func (\*SipStack) [CreateResponseChannel](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L51)

```go
func (ed *SipStack) CreateResponseChannel() chan transaction.Transaction
```

### func (\*SipStack) [GetTransactionManager](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L33)

```go
func (ed *SipStack) GetTransactionManager() *transaction.TransactionManager
```

### func (\*SipStack) [IsAlive](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L58)

```go
func (ed *SipStack) IsAlive() bool
```

### func (\*SipStack) [Start](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L68)

```go
func (ed *SipStack) Start()
```

Start starts the sip stack

### func (\*SipStack) [Stop](https://github.com/KalbiProject/kalbi/blob/master/kalbi.go#L62)

```go
func (ed *SipStack) Stop()
```

# sip

```go
import "github.com/KalbiProject/kalbi/sip"
```

## func [Parse](https://github.com/KalbiProject/kalbi/blob/master/sip/sip.go#L8)

```go
func Parse(msg []byte) message.SipMsg
```

# message

```go
import "github.com/KalbiProject/kalbi/sip/message"
```

## Constants

```go
const fieldAddrType = 40
```

```go
const fieldBase = 1
```

```go
const fieldBranch = 12
```

```go
const fieldCat = 45
```

```go
const fieldConnAddr = 41
```

```go
const fieldExpires = 17
```

```go
const fieldFmt = 44
```

```go
const fieldUserHost = 6
```

```go
const fieldID = 9
```

```go
const fieldIgnore = 255
```

```go
const fieldMaddr = 14
```

```go
const fieldMedia = 42
```

```go
const fieldMethod = 10
```

```go
const fieldName = 3
```

```go
const fieldNameQ = 4
```

```go
const fieldNull = 0
```

```go
const fieldPort = 7
```

```go
const fieldProto = 43
```

```go
const fieldQ = 18
```

```go
const fieldRec = 16
```

```go
const fieldRport = 13
```

```go
const fieldStatus = 20
```

```go
const fieldStatusDesc = 21
```

```go
const fieldTag = 8
```

```go
const fieldTran = 11
```

```go
const fieldTTL = 15
```

```go
const fieldUser = 5
```

```go
const fieldUserType = 19
```

```go
const fieldValue = 2
```

## func [GenerateBranchId](https://github.com/KalbiProject/kalbi/blob/master/sip/message/utils.go#L9)

```go
func GenerateBranchId() string
```

## func [MessageDetails](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L212)

```go
func MessageDetails(data *SipMsg) string
```

MessageDetails prints all we know about the struct in a readable format

## func [ParseSipContact](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipContact.go#L62)

```go
func ParseSipContact(v []byte, out *SipContact)
```

## func [ParseSipCseq](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipCseq.go#L36)

```go
func ParseSipCseq(v []byte, out *SipCseq)
```

## func [ParseSipFrom](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L55)

```go
func ParseSipFrom(v []byte, out *SipFrom)
```

## func [ParseSipReq](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L71)

```go
func ParseSipReq(v []byte, out *SipReq)
```

## func [ParseSipTo](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L51)

```go
func ParseSipTo(v []byte, out *SipTo)
```

## func [ParseSipVia](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipVia.go#L42)

```go
func ParseSipVia(v []byte, out *SipVia)
```

## type [SipContact](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipContact.go#L22-L32)

```go
type SipContact struct {
    UriType string // Type of URI sip, sips, tel etc
    Name    []byte // Named portion of URI
    User    []byte // User part
    Host    []byte // Host part
    Port    []byte // Port number
    Tran    []byte // Transport
    Qval    []byte // Q Value
    Expires []byte // Expires
    Src     []byte // Full source if needed
}
```

### func (\*SipContact) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipContact.go#L39)

```go
func (sc *SipContact) String() string
```

### func (\*SipContact) [SetName](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipContact.go#L34)

```go
func (sc *SipContact) SetName(name string)
```

## type [SipCseq](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipCseq.go#L16-L20)

```go
type SipCseq struct {
    ID     []byte // Cseq ID
    Method []byte // Cseq Method
    Src    []byte // Full source if needed
}
```

### func (\*SipCseq) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipCseq.go#L30)

```go
func (sc *SipCseq) String() string
```

### func (\*SipCseq) [SetID](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipCseq.go#L22)

```go
func (sc *SipCseq) SetID(id string)
```

### func (\*SipCseq) [SetMethod](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipCseq.go#L26)

```go
func (sc *SipCseq) SetMethod(method string)
```

## type [SipFrom](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L11-L20)

```go
type SipFrom struct {
    UriType  string // Type of URI sip, sips, tel etc
    Name     []byte // Named portion of URI
    User     []byte // User part
    Host     []byte // Host part
    Port     []byte // Port number
    Tag      []byte // Tag
    UserType []byte // User Type
    Src      []byte // Full source if needed
}
```

### func (\*SipFrom) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L46)

```go
func (sf *SipFrom) String() string
```

### func (\*SipFrom) [SetHost](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L30)

```go
func (sf *SipFrom) SetHost(host string)
```

### func (\*SipFrom) [SetPort](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L34)

```go
func (sf *SipFrom) SetPort(port string)
```

### func (\*SipFrom) [SetTag](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L42)

```go
func (sf *SipFrom) SetTag(tag string)
```

### func (\*SipFrom) [SetUriType](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L22)

```go
func (sf *SipFrom) SetUriType(uriType string)
```

### func (\*SipFrom) [SetUser](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L26)

```go
func (sf *SipFrom) SetUser(user string)
```

### func (\*SipFrom) [SetUserType](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipFrom.go#L38)

```go
func (sf *SipFrom) SetUserType(userType string)
```

## type [SipMsg](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L14-L30)

SipMsg is a representation of a SIP message

```go
type SipMsg struct {
    Req      SipReq
    From     SipFrom
    To       SipTo
    Contact  SipContact
    Via      []SipVia
    Cseq     SipCseq
    Ua       SipVal
    Exp      SipVal
    MaxFwd   SipVal
    CallID   SipVal
    ContType SipVal
    ContLen  SipVal
    Src      []byte

    Sdp sdp.SdpMsg
}
```

### func [NewRequest](https://github.com/KalbiProject/kalbi/blob/master/sip/message/request.go#L8)

```go
func NewRequest(request string, to string, from string) *SipMsg
```

NewRequest creates new SIP request

### func [NewResponse](https://github.com/KalbiProject/kalbi/blob/master/sip/message/response.go#L9)

```go
func NewResponse(response int, to string, from string) *SipMsg
```

NewResponse creates new SIP response

### func [Parse](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L98)

```go
func Parse(v []byte) (output SipMsg)
```

Parse routine\, passes by value

### func (\*SipMsg) [CopyHeaders](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L42)

```go
func (sm *SipMsg) CopyHeaders(msg *SipMsg)
```

CopyHeaders copys headers from one SIP message to another

### func (\*SipMsg) [CopySdp](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L55)

```go
func (sm *SipMsg) CopySdp(msg *SipMsg)
```

CopySdp copys SDP from one SIP message to another

### func (\*SipMsg) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L60)

```go
func (sm *SipMsg) String() string
```

Export returns SIP message as string

### func (\*SipMsg) [GetStatusCode](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L33)

```go
func (sm *SipMsg) GetStatusCode() int
```

GetStatusCode returns responses status code

## type [SipReq](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L15-L25)

```go
type SipReq struct {
    Method     []byte // Sip Method eg INVITE etc
    UriType    string // Type of URI sip, sips, tel etc
    StatusCode []byte // Status Code eg 100
    StatusDesc []byte // Status Code Description eg trying
    User       []byte // User part
    Host       []byte // Host part
    Port       []byte // Port number
    UserType   []byte // User Type
    Src        []byte // Full source if needed
}
```

### func (\*SipReq) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L59)

```go
func (sr *SipReq) String() string
```

### func (\*SipReq) [SetHost](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L47)

```go
func (sr *SipReq) SetHost(host string)
```

### func (\*SipReq) [SetMethod](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L27)

```go
func (sr *SipReq) SetMethod(method string)
```

### func (\*SipReq) [SetPort](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L51)

```go
func (sr *SipReq) SetPort(port string)
```

### func (\*SipReq) [SetStatusCode](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L35)

```go
func (sr *SipReq) SetStatusCode(code int)
```

### func (\*SipReq) [SetStatusDesc](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L39)

```go
func (sr *SipReq) SetStatusDesc(desc string)
```

### func (\*SipReq) [SetUriType](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L31)

```go
func (sr *SipReq) SetUriType(uriType string)
```

### func (\*SipReq) [SetUser](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L43)

```go
func (sr *SipReq) SetUser(user string)
```

### func (\*SipReq) [SetUserType](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipRequestLine.go#L55)

```go
func (sr *SipReq) SetUserType(userType string)
```

## type [SipTo](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L7-L16)

```go
type SipTo struct {
    UriType  string // Type of URI sip, sips, tel etc
    Name     []byte // Named portion of URI
    User     []byte // User part
    Host     []byte // Host part
    Port     []byte // Port number
    Tag      []byte // Tag
    UserType []byte // User Type
    Src      []byte // Full source if needed
}
```

### func (\*SipTo) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L42)

```go
func (sf *SipTo) String() string
```

### func (\*SipTo) [SetHost](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L26)

```go
func (sf *SipTo) SetHost(host string)
```

### func (\*SipTo) [SetPort](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L30)

```go
func (sf *SipTo) SetPort(port string)
```

### func (\*SipTo) [SetTag](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L38)

```go
func (sf *SipTo) SetTag(tag string)
```

### func (\*SipTo) [SetUriType](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L18)

```go
func (sf *SipTo) SetUriType(uriType string)
```

### func (\*SipTo) [SetUser](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L22)

```go
func (sf *SipTo) SetUser(user string)
```

### func (\*SipTo) [SetUserType](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipTo.go#L34)

```go
func (sf *SipTo) SetUserType(userType string)
```

## type [SipVal](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L83-L86)

```go
type SipVal struct {
    Value []byte // Sip Value
    Src   []byte // Full source if needed
}
```

### func (\*SipVal) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L93)

```go
func (sv *SipVal) String() string
```

Export returns SIP value as string

### func (\*SipVal) [SetValue](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sip.go#L88)

```go
func (sv *SipVal) SetValue(value string)
```

## type [SipVia](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipVia.go#L14-L24)

```go
type SipVia struct {
    Trans  string // Type of Transport udp, tcp, tls, sctp etc
    Host   []byte // Host part
    Port   []byte // Port number
    Branch []byte //
    Rport  []byte //
    Maddr  []byte //
    Ttl    []byte //
    Rcvd   []byte //
    Src    []byte // Full source if needed
}
```

### func (\*SipVia) [Export](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipVia.go#L26)

```go
func (sv *SipVia) String() string
```

### func (\*SipVia) [SetBranch](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipVia.go#L38)

```go
func (sv *SipVia) SetBranch(value string)
```

### func (\*SipVia) [SetHost](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipVia.go#L30)

```go
func (sv *SipVia) SetHost(value string)
```

### func (\*SipVia) [SetPort](https://github.com/KalbiProject/kalbi/blob/master/sip/message/sipVia.go#L34)

```go
func (sv *SipVia) SetPort(value string)
```

Generated by [gomarkdoc](https://github.com/princjef/gomarkdoc)

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# method

```go
import "github.com/KalbiProject/kalbi/sip/method"
```

## Constants

```go
const (
    //REGISTER Register the URI listed in the To-header field with a location server and associates it with the network address given in a Contact header field. RFC_3261
    REGISTER = "REGISTER"
    //INVITE Initiate a dialog for establishing a call. The request is sent by a user agent client to a user agent server.
    INVITE = "INVITE"
    //ACK Confirm that an entity has received a final response to an INVITE request.
    ACK = "ACK"
    //BYE Signal termination of a dialog and end a call.
    BYE = "BYE"
    //CANCEL Cancel any pending request.
    CANCEL = "CANCEL"
    //UPDATE Modify the state of a session without changing the state of the dialog.
    UPDATE = "UPDATE"
    //REFER Ask recipient to issue a request for the purpose of call transfer.
    REFER = "REFER"
    //PRACK Provisional acknowledgement.
    PRACK = "PRACK"
    //SUBSCRIBE Initiates a subscription for notification of events from a notifier.
    SUBSCRIBE = "SUBSCRIBE"
    //NOTIFY Inform a subscriber of notifications of a new event.
    NOTIFY = "NOTIFY"
    //PUBLISH Publish an event to a notification server.
    PUBLISH = "PUBLISH"
    //MESSAGE Deliver a text message.
    MESSAGE = "MESSAGE"
    //INFO Send mid-session information that does not modify the session state.
    INFO = "INFO"
    //OPTIONS Query the capabilities of an endpoint.
    OPTIONS = "OPTIONS"
)
```

# status

```go
import "github.com/KalbiProject/kalbi/sip/status"
```

## Constants

```go
const (

    // Trying Extended search being performed may take a significant time so a forking proxy must send a 100 Trying response.
    // RFC_3261
    Trying = 100 // : "Trying"
    // Ringing Destination user agent received INVITE, and is alerting user of call.
    // RFC_3261
    Ringing = 180 // : "Ringing",
    // CallIsBeingForwarded Servers can optionally send this response to indicate a call is being forwarded.
    // RFC_3261
    CallIsBeingForwarded = 181 // : "CallIsBeingForwarded",
    // Queued Indicates that the destination was temporarily unavailable, so the server has queued the call until the destination is available.
    // A server may send multiple 182 responses to update progress of the queue.
    // RFC_3261
    Queued = 182 // : "Queued",
    // SessionProgress This response may be used to send extra information for a call which is still being set up.
    // RFC_3261
    SessionProgress = 183 // : "SessionProgress",
    // EarlyDialogTerminated Can be used by User Agent Server to indicate to upstream SIP entities (including the User Agent Client (UAC)) that an early dialog has been terminated.
    // RFC_6228
    EarlyDialogTerminated = 199 // : "EarlyDialogTerminated",

    // OK Indicates that the request was successful.
    // RFC_3261
    OK  = 200 // : "OK",
    // Accepted Indicates that the request has been accepted for processing, but the processing has not been completed.
    // RFC_3265
    // RFC_2616 10.2.3
    // Deprecated RFC_6665
    Accepted = 202 // : "Accepted",
    // NoNotification Indicates the request was successful, but the corresponding response will not be received.
    // RFC_5839 7.1
    NoNotification = 204 // : "NoNotification",

    // MultipleChoices The address resolved to one of several options for the user or client to choose between,
    // which are listed in the message body or the message's Contact fields.
    // RFC_3261
    MultipleChoices = 300 // : "MultipleChoices",
    // MovedPermanently The original Request-URI is no longer valid, the new address is given in the Contact header field,
    // and the client should update any records of the original Request-URI with the new value.
    // RFC_3261
    MovedPermanently = 301 // : "MovedPermanently",
    // MovedTemporarily The client should try at the address in the Contact field.
    // If an Expires field is present, the client may cache the result for that period of time.
    // RFC_3261
    MovedTemporarily = 302 // : "MovedTemporarily",
    // UseProxy The Contact field details a proxy that must be used to access the requested destination.
    // RFC_3261
    UseProxy = 305 // : "UseProxy",
    // AlternativeService The call failed, but alternatives are detailed in the message body.
    // RFC_3261
    AlternativeService = 380 // : "AlternativeService",

    // BadRequest The request could not be understood due to malformed syntax.
    // RFC_3261
    BadRequest = 400 // : "BadRequest",
    // Unauthorized The request requires user authentication. This response is issued by UASs and registrars.
    // RFC_3261
    Unauthorized = 401 // : "Unauthorized",
    // PaymentRequired Reserved for future use.
    // RFC_3261
    PaymentRequired = 402 // : "PaymentRequired",
    // Forbidden The server understood the request, but is refusing to fulfill it.
    // RFC_3261
    // Sometimes (but not always) this means the call has been rejected by the receiver.
    Forbidden = 403 // : "Forbidden",
    // NotFound The server has definitive information that the user does not exist at the domain specified in the Request-URI.
    // This status is also returned if the domain in the Request-URI does not match any of the domains handled by the recipient of the request.
    // RFC_3261
    NotFound = 404 // : "NotFound",
    // MethodNotAllowed The method specified in the Request-Line is understood, but not allowed for the address identified by the Request-URI.
    // RFC_3261
    MethodNotAllowed = 405 // : "MethodNotAllowed",
    // NotAcceptable406 The resource identified by the request is only capable of generating response entities that have content characteristics but
    // not acceptable according to the Accept header field sent in the request.
    // RFC_3261
    NotAcceptable406 = 406 // : "NotAcceptable406",
    // ProxyAuthenticationRequired The request requires user authentication. This response is issued by proxys.
    // RFC_3261
    ProxyAuthenticationRequired = 407 // : "ProxyAuthenticationRequired",
    // RequestTimeout Couldn't find the user in time.
    // The server could not produce a response within a suitable amount of time,
    // for example, if it could not determine the location of the user in time.
    // The client MAY repeat the request without modifications at any later time.
    // RFC_3261
    RequestTimeout = 408 // : "RequestTimeout",
    // Conflict User already registered.
    // RFC_2543
    // Deprecated by omission from later RFCs RFC_3261 and by non-registration with the IANA.
    Conflict = 409 // : "Conflict",
    // Gone The user existed once, but is not available here any more.
    // RFC_3261
    Gone = 410 // : "Gone",
    // LengthRequired The server will not accept the request without a valid Content-Length. RFC_2543
    // Deprecated by omission from later RFCs RFC_3261 and by non-registration with the IANA.
    LengthRequired = 411 // : "LengthRequired",
    // ConditionalRequestFailed The given precondition has not been met.
    // RFC_3903 11.2.1
    ConditionalRequestFailed = 412 // : "ConditionalRequestFailed",
    // RequestEntityTooLarge Request body too large.
    // RFC_3261
    RequestEntityTooLarge = 413 // : "RequestEntityTooLarge",
    // RequestURITooLong The server is refusing to service the request because the Request-URI is longer than the server is willing to interpret.
    // RFC_3261
    RequestURITooLong = 414 // : "RequestURITooLong",
    // UnsupportedMediaType Request body in a format not supported.
    // RFC_3261
    UnsupportedMediaType = 415 // : "UnsupportedMediaType",
    // UnsupportedURIScheme Request-URI is unknown to the server.
    // RFC_3261
    UnsupportedURIScheme = 416 // : "UnsupportedURIScheme",
    // UnknownResourcePriority There was a resource-priority option tag, but no Resource-Priority header.
    // RFC_4412 4.6.2
    UnknownResourcePriority = 417 // : "UnknownResourcePriority",
    // BadExtension Bad SIP Protocol Extension used, not understood by the server.
    // RFC_3261
    BadExtension = 420 // : "BadExtension",
    // ExtensionRequired The server needs a specific extension not listed in the Supported header.
    // RFC_3261
    ExtensionRequired = 421 // : "ExtensionRequired",
    // SessionIntervalTooSmall The received request contains a Session-Expires header field with a duration below the minimum timer.
    // RFC_4028_sec_6
    SessionIntervalTooSmall = 422 // : "SessionIntervalTooSmall",
    // IntervalTooBrief Expiration time of the resource is too short.
    // RFC_3261
    IntervalTooBrief = 423 // : "IntervalTooBrief",
    // BadLocationInformation The request's location content was malformed or otherwise unsatisfactory.
    // RFC_6422
    // Location Conveyance for the Session Initiation Protocol RFC_6442 4.3
    BadLocationInformation = 424 // : "BadLocationInformation",
    // UseIdentityHeader The server policy requires an Identity header, and one has not been provided.
    // RFC_4474
    UseIdentityHeader = 428 // : "UseIdentityHeader",
    // ProvideReferrerIdentity The server did not receive a valid Referred-By token on the request.
    // RFC_3892 5
    ProvideReferrerIdentity = 429 // : "ProvideReferrerIdentity",
    // FlowFailed A specific flow to a user agent has failed, although other flows may succeed.
    // This response is intended for use between proxy devices,
    // and should not be seen by an endpoint (and if it is seen by one, should be treated as a [[#400|400 Bad Request]] response).
    // RFC_5626
    FlowFailed = 430 // : "FlowFailed",
    // AnonymityDisallowed The request has been rejected because it was anonymous.
    // RFC_5079 5
    AnonymityDisallowed = 433 // : "AnonymityDisallowed",
    // BadIdentityInfo The request has an Identity-Info header, and the URI scheme in that header cannot be dereferenced.
    // RFC_4474
    BadIdentityInfo = 436 // : "BadIdentityInfo",
    // UnsupportedCertificate The server was unable to validate a certificate for the domain that signed the request.<ref name="RFC_4474"/>{{rp|p11}}
    UnsupportedCertificate = 437 // : "UnsupportedCertificate",
    // InvalidIdentityHeader The server obtained a valid certificate that the request claimed was used to sign the request, but was unable to verify that signature.<ref name="RFC_4474"/>{{rp|p12}}
    InvalidIdentityHeader = 438 // : "InvalidIdentityHeader",
    // FirstHopLacksOutboundSupport The first outbound [[Session Initiation Protocol#Proxy server|proxy]] the user is attempting to register through does not
    // support the "outbound" feature of RFC 5626, although the [[Session Initiation Protocol#Registrar|registrar]] does.<ref name="RFC_5626"/>{{rp|§11.6}}
    FirstHopLacksOutboundSupport = 439 // : "FirstHopLacksOutboundSupport",
    // MaxBreadthExceeded If a SIP proxy determines a response context has insufficient Incoming Max-Breadth to carry out a desired parallel fork,
    // and the proxy is unwilling/unable to compensate by forking serially or sending a redirect, that proxy MUST return a 440 response.
    // A client receiving a 440 response can infer that its request did not reach all possible destinations.
    // RFC_5393
    MaxBreadthExceeded = 440 // : "MaxBreadthExceeded",
    // BadInfoPackage If a SIP UA receives an INFO request associated with an Info Package that the UA has not indicated willingness to receive,
    // the UA MUST send a 469 response, which contains a Recv-Info header field with Info Packages for which the UA is willing to receive INFO requests.
    // 6086
    BadInfoPackage = 469 // : "BadInfoPackage",
    // ConsentNeeded The source of the request did not have the permission of the recipient to make such a request.<ref name="RFC_5360">{{cite IETF
    //| section      = 5.9.2
    ConsentNeeded = 470 // : "ConsentNeeded",
    // TemporarilyUnavailable Callee currently unavailable.<ref name="RFC_3261"/>{{rp|§21.4.18}}
    TemporarilyUnavailable = 480 // : "TemporarilyUnavailable",
    // TransactionDoesNotExist Server received a request that does not match any dialog or transaction.<ref name="RFC_3261"/>{{rp|§21.4.19}}
    TransactionDoesNotExist = 481 // : "TransactionDoesNotExist",
    // LoopDetected Server has detected a loop.<ref name="RFC_3261"/>{{rp|§21.4.20}}
    LoopDetected = 482 // : "LoopDetected",
    // TooManyHops Max-Forwards header has reached the value '0'.<ref name="RFC_3261"/>{{rp|§21.4.21}}
    TooManyHops = 483 // : "TooManyHops",
    // AddressIncomplete Request-URI incomplete.<ref name="RFC_3261"/>{{rp|§21.4.22}}
    AddressIncomplete = 484 // : "AddressIncomplete",
    // Ambiguous Request-URI is ambiguous.<ref name="RFC_3261"/>{{rp|§21.4.23}}
    Ambiguous = 485 // : "Ambiguous",
    // BusyHere Callee is busy.<ref name="RFC_3261"/>{{rp|§21.4.24}}
    BusyHere = 486 // : "BusyHere",
    // RequestTerminated Request has terminated by bye or cancel.<ref name="RFC_3261"/>{{rp|§21.4.25}}
    RequestTerminated = 487 // : "RequestTerminated",
    // NotAcceptableHere Some aspect of the session description or the Request-URI is not acceptable.<ref name="RFC_3261"/>{{rp|§21.4.26}}
    NotAcceptableHere = 488 // : "NotAcceptableHere",
    // BadEvent The server did not understand an event package specified in an Event header field.<ref name="RFC_3265"/>{{rp|§7.3.2}}<ref name="RFC_6665"/>{{rp|§8.3.2}}
    BadEvent = 489 // : "BadEvent",
    // RequestPending Server has some pending request from the same dialog.<ref name="RFC_3261"/>{{rp|§21.4.27}}
    RequestPending = 491 // : "RequestPending",
    // Undecipherable Request contains an encrypted MIME body, which recipient can not decrypt.<ref name="RFC_3261"/>{{rp|§21.4.28}}
    Undecipherable = 493 // : "Undecipherable",
    // SecurityAgreementRequired The server has received a request that requires a negotiated security mechanism,
    // and the response contains a list of suitable security mechanisms for the requester to choose between, RFC_3329 or a [[digest authentication]] challenge.
    // RFC_3329
    SecurityAgreementRequired = 494 // : "SecurityAgreementRequired",

    // InternalServerError The server could not fulfill the request due to some unexpected condition.<ref name="RFC_3261"/>{{rp|§21.5.1}}
    InternalServerError = 500 // : "InternalServerError",
    // NotImplemented The server does not have the ability to fulfill the request, such as because it does not recognize the request method.
    // (Compare with [[#405|405 Method Not Allowed]], where the server recognizes the method but does not allow or support it.)
    // RFC_3261
    NotImplemented = 501 // : "NotImplemented",
    // BadGateway The server is acting as a [[Gateway (computer program)|gateway]] or [[Proxy server|proxy]], and received an invalid response from a downstream server while attempting to fulfill the request.<ref name="RFC_3261"/>{{rp|§21.5.3}}
    BadGateway = 502 // : "BadGateway",
    // ServiceUnavailable The server is undergoing maintenance or is temporarily overloaded and so cannot process the request. A "Retry-After" header field may specify when the client may reattempt its request.<ref name="RFC_3261"/>{{rp|§21.5.4}}
    ServiceUnavailable = 503 // : "ServiceUnavailable",
    // ServerTimeout The server attempted to access another server in attempting to process the request, and did not receive a prompt response.<ref name="RFC_3261"/>{{rp|§21.5.5}}
    ServerTimeout = 504 // : "ServerTimeout",
    // VersionNotSupported The SIP protocol version in the request is not supported by the server.<ref name="RFC_3261"/>{{rp|§21.5.6}}
    VersionNotSupported = 505 // : "VersionNotSupported",
    // MessageTooLarge The request message length is longer than the server can process.<ref name="RFC_3261"/>{{rp|§21.5.7}}
    MessageTooLarge = 513 // : "MessageTooLarge",
    // PushNotificationServiceNotSupported The server does not support the push notification service identified in a 'pn-provider' SIP URI parameter<ref name="RFC_8599">{{cite IETF
    PushNotificationServiceNotSupported = 555 // : "PushNotificationServiceNotSupported",
    // PreconditionFailure The server is unable or unwilling to meet some constraints specified in the offer.
    // RFC_3312 8
    PreconditionFailure = 580 // : "PreconditionFailure",

    // BusyEverywhere All possible destinations are busy. Unlike the 486 response, this response indicates the destination knows there are no alternative destinations (such as a voicemail server) able to accept the call.<ref name="RFC_3261"/>{{rp|§21.6.1}}
    BusyEverywhere = 600 // : "BusyEverywhere",
    // Decline The destination does not wish to participate in the call, or cannot do so, and additionally the destination knows there are no alternative destinations (such as a voicemail server) willing to accept the call.<ref name="RFC_3261"/>{{rp|§21.6.2}} The response may indicate a better time to call in the Retry-After header field.
    Decline = 603 // : "Decline",
    // DoesNotExistAnywhere The server has authoritative information that the requested user does not exist anywhere.<ref name="RFC_3261"/>{{rp|§21.6.3}}
    DoesNotExistAnywhere = 604 // : "DoesNotExistAnywhere",
    // NotAcceptable606 The user's agent was contacted successfully but some aspects of the session description such as the requested media, bandwidth, or addressing style were not acceptable.<ref name="RFC_3261"/>{{rp|§21.6.4}}
    NotAcceptable606 = 606 // : "NotAcceptable606",
    // Unwanted The called party did not want this call from the calling party. Future attempts from the calling party are likely to be similarly rejected.<ref name="RFC_8197">{{cite IETF
    Unwanted = 607 // : "Unwanted",
    // Rejected An intermediary machine or process rejected the call attempt.
    // RFC_8688
    // This contrasts with the 607 (Unwanted) SIP response code in which the called party rejected the call.
    // The response may include the contact entities that blocked the call in Call-Info header containing.
    // This provides a remediation mechanism for legal callers that find their calls blocked.
    Rejected = 608 // : "Rejected",
)
```

## func [StatusText](https://github.com/KalbiProject/kalbi/blob/master/sip/status/status.go#L336)

```go
func StatusText(code int) string
```

StatusText returns a text for the HTTP status code\. It returns the empty string if the code is unknown\.

# transaction

```go
import "github.com/KalbiProject/kalbi/sip/transaction"
```

## Constants

```go
const (
    //T1 is a timer described in RFC3261.
    T1  = 500 * time.Millisecond
    //T2 is a timer described in RFC3261.
    T2  = 4 * time.Second
)
```

## func [GenerateBranchId](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/utils.go#L9)

```go
func GenerateBranchId() string
```

## type [ClientTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L61-L78)

ClientTransaction represents a client transaction refrences in RFC3261

```go
type ClientTransaction struct {
    ID           string
    BranchID     string
    ServerTxID   string
    TransManager *TransactionManager
    Origin       *message.SipMsg
    FSM          *fsm.FSM

    ListeningPoint transport.ListeningPoint
    Host           string
    Port           string
    LastMessage    *message.SipMsg
    // contains filtered or unexported fields
}
```

### func (\*ClientTransaction) [GetBranchID](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L114)

```go
func (ct *ClientTransaction) GetBranchID() string
```

GetBranchID returns branchId which is the identifier of a transaction

### func (\*ClientTransaction) [GetLastMessage](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L147)

```go
func (ct *ClientTransaction) GetLastMessage() *message.SipMsg
```

GetLastMessage returns the last received SIP message to this transaction

### func (\*ClientTransaction) [GetOrigin](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L119)

```go
func (ct *ClientTransaction) GetOrigin() *message.SipMsg
```

GetOrigin returns the SIP message that initiated this transaction

### func (\*ClientTransaction) [GetServerTransactionID](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L142)

```go
func (ct *ClientTransaction) GetServerTransactionID() string
```

GetServerTransaction returns a ServerTransaction that has been set with SetServerTransaction\(\)

### func (\*ClientTransaction) [InitFSM](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L81)

```go
func (ct *ClientTransaction) InitFSM(msg *message.SipMsg)
```

InitFSM initializes the finite state machine within the client transaction

### func (\*ClientTransaction) [Receive](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L124)

```go
func (ct *ClientTransaction) Receive(msg *message.SipMsg)
```

Receive takes in the SIP message from the transport layer

### func (\*ClientTransaction) [Resend](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L189)

```go
func (ct *ClientTransaction) Resend()
```

Resend is used for retransmissions

### func (\*ClientTransaction) [Send](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L207)

```go
func (ct *ClientTransaction) Send(msg *message.SipMsg, host string, port string)
```

Send is used to send a SIP message

### func (\*ClientTransaction) [SetLastMessage](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L152)

```go
func (ct *ClientTransaction) SetLastMessage(msg *message.SipMsg)
```

SetLastMessage sets the last message received

### func (\*ClientTransaction) [SetListeningPoint](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L109)

```go
func (ct *ClientTransaction) SetListeningPoint(lp transport.ListeningPoint)
```

SetListeningPoint sets a listening point to the client transaction

### func (\*ClientTransaction) [SetServerTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L137)

```go
func (ct *ClientTransaction) SetServerTransaction(txID string)
```

SetServerTransaction is used to set a Server Transaction

### func (\*ClientTransaction) [StatelessSend](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/client.go#L197)

```go
func (ct *ClientTransaction) StatelessSend(msg *message.SipMsg, host string, port string)
```

StatelessSend send a sip message without acting on the FSM

## type [ServerTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L45-L56)

ServerTransaction is a representation of a Server Transaction refrences in RFC3261

```go
type ServerTransaction struct {
    ID           string
    BranchID     string
    TransManager *TransactionManager
    Origin       *message.SipMsg
    FSM          *fsm.FSM

    ListeningPoint transport.ListeningPoint
    Host           string
    Port           string
    LastMessage    *message.SipMsg
    // contains filtered or unexported fields
}
```

### func (\*ServerTransaction) [GetBranchID](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L93)

```go
func (st *ServerTransaction) GetBranchID() string
```

GetBranchID returns branchId which is the identifier of a transaction

### func (\*ServerTransaction) [GetLastMessage](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L103)

```go
func (st *ServerTransaction) GetLastMessage() *message.SipMsg
```

GetLastMessage returns the last received SIP message to this transaction

### func (\*ServerTransaction) [GetOrigin](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L98)

```go
func (st *ServerTransaction) GetOrigin() *message.SipMsg
```

GetOrigin returns the SIP message that initiated this transaction

### func (\*ServerTransaction) [GetServerTransactionID](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L138)

```go
func (st *ServerTransaction) GetServerTransactionID() string
```

GetServerTransactionID returns Server transaction ID

### func (\*ServerTransaction) [InitFSM](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L59)

```go
func (st *ServerTransaction) InitFSM(msg *message.SipMsg)
```

InitFSM initializes the finite state machine within the client transaction

### func (\*ServerTransaction) [Receive](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L113)

```go
func (st *ServerTransaction) Receive(msg *message.SipMsg)
```

Receive takes in the SIP message from the transport layer

### func (\*ServerTransaction) [Respond](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L124)

```go
func (st *ServerTransaction) Respond(msg *message.SipMsg)
```

Respond is used to process response from transport layer

### func (\*ServerTransaction) [Send](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L143)

```go
func (st *ServerTransaction) Send(msg *message.SipMsg, host string, port string)
```

Send used to send SIP message to specified host

### func (\*ServerTransaction) [SetLastMessage](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L108)

```go
func (st *ServerTransaction) SetLastMessage(msg *message.SipMsg)
```

SetLastMessage sets the last message received

### func (\*ServerTransaction) [SetListeningPoint](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/server.go#L88)

```go
func (st *ServerTransaction) SetListeningPoint(lp transport.ListeningPoint)
```

SetListeningPoint sets a listening point to the client transaction

## type [Transaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/transaction.go#L109-L118)

Transaction

```go
type Transaction interface {
    GetBranchID() string
    GetOrigin() *message.SipMsg
    SetListeningPoint(transport.ListeningPoint)
    Send(*message.SipMsg, string, string)
    Receive(*message.SipMsg)
    GetLastMessage() *message.SipMsg
    GetServerTransactionID() string
    SetLastMessage(*message.SipMsg)
}
```

## type [TransactionManager](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L23-L30)

TransactionManager handles SIP transactions

```go
type TransactionManager struct {
    ServerTX        map[string]Transaction
    ClientTX        map[string]Transaction
    RequestChannel  chan Transaction
    ResponseChannel chan Transaction
    ListeningPoint  transport.ListeningPoint
    // contains filtered or unexported fields
}
```

### func [NewTransactionManager](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L13)

```go
func NewTransactionManager() *TransactionManager
```

NewTransactionManager returns a new TransactionManager

### func (\*TransactionManager) [DeleteClientTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L128)

```go
func (tm *TransactionManager) DeleteClientTransaction(branch string)
```

DeleteClientTransaction removes a stored transaction

### func (\*TransactionManager) [DeleteServerTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L118)

```go
func (tm *TransactionManager) DeleteServerTransaction(branch string)
```

DeleteServerTransaction removes a stored transaction

### func (\*TransactionManager) [FindClientTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L77)

```go
func (tm *TransactionManager) FindClientTransaction(msg *message.SipMsg) (Transaction, bool)
```

FindTransaction finds transaction by SipMsgg struct

### func (\*TransactionManager) [FindClientTransactionByID](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L95)

```go
func (tm *TransactionManager) FindClientTransactionByID(value string) (Transaction, bool)
```

FindTransactionByID finds transaction by id

### func (\*TransactionManager) [FindServerTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L68)

```go
func (tm *TransactionManager) FindServerTransaction(msg *message.SipMsg) (Transaction, bool)
```

FindTransaction finds transaction by SipMsgg struct

### func (\*TransactionManager) [FindServerTransactionByID](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L86)

```go
func (tm *TransactionManager) FindServerTransactionByID(value string) (Transaction, bool)
```

FindTransactionByID finds transaction by id

### func (\*TransactionManager) [Handle](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L33)

```go
func (tm *TransactionManager) Handle(message *message.SipMsg)
```

Handle runs TransManager

### func (\*TransactionManager) [MakeKey](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L138)

```go
func (tm *TransactionManager) MakeKey(msg message.SipMsg) string
```

MakeKey creates new transaction identifier

### func (\*TransactionManager) [NewClientTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L157)

```go
func (tm *TransactionManager) NewClientTransaction(msg *message.SipMsg) *ClientTransaction
```

NewClientTransaction builds new CLientTransaction

### func (\*TransactionManager) [NewServerTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L174)

```go
func (tm *TransactionManager) NewServerTransaction(msg *message.SipMsg) *ServerTransaction
```

NewServerTransaction builds new ServerTransaction

### func (\*TransactionManager) [PutClientTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L111)

```go
func (tm *TransactionManager) PutClientTransaction(tx Transaction)
```

PutClientTransaction stores a transaction

### func (\*TransactionManager) [PutServerTransaction](https://github.com/KalbiProject/kalbi/blob/master/sip/transaction/manager.go#L104)

```go
func (tm *TransactionManager) PutServerTransaction(tx Transaction)
```

PutServerTransaction stores a transaction

# sdp

```go
import "github.com/KalbiProject/kalbi/sdp"
```

## Constants

```go
const fieldAddrType = 4
```

```go
const fieldBase = 1
```

```go
const fieldCat = 9
```

```go
const fieldConnAddr = 5
```

```go
const fieldFmt = 8
```

```go
const fieldIgnore = 255
```

```go
const fieldMedia = 6
```

```go
const fieldNetType = 13
```

FMS States

```go
const fieldNull = 0
```

```go
const fieldPort = 3
```

```go
const fieldProto = 7
```

```go
const fieldSessionID = 11
```

```go
const fieldSessionVersion = 12
```

```go
const fieldTimeStart = 15
```

```go
const fieldTimeStop = 16
```

```go
const fieldUniAddr = 14
```

```go
const fieldUsername = 10
```

```go
const fieldValue = 2
```

## func [ParseSdpOrigin](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdpOrigin.go#L90)

```go
func ParseSdpOrigin(v []byte, out *SdpOrigin)
```

## func [ParserSdpTime](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdpTime.go#L69)

```go
func ParserSdpTime(v []byte, out *sdpTime)
```

## type [SdpMsg](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdp.go#L32-L39)

SdpMsg is representation of an SDP message

```go
type SdpMsg struct {
    Origin    SdpOrigin
    Version   sdpVersion
    Time      sdpTime
    MediaDesc sdpMediaDesc
    Attrib    []sdpAttrib
    ConnData  sdpConnData
}
```

### func [Parse](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdp.go#L72)

```go
func Parse(v []byte) (output SdpMsg)
```

Parse Parses SDP

### func (\*SdpMsg) [Export](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdp.go#L47)

```go
func (sm *SdpMsg) String() string
```

### func (\*SdpMsg) [Size](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdp.go#L42)

```go
func (sm *SdpMsg) Size() int
```

Size returns size in bytes

## type [SdpOrigin](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdpOrigin.go#L69-L77)

```go
type SdpOrigin struct {
    Username       []byte
    SessionId      []byte
    SessionVersion []byte
    NetType        []byte
    AddrType       []byte
    UniAddr        []byte
    Src            []byte
}
```

### func (\*SdpOrigin) [Export](https://github.com/KalbiProject/kalbi/blob/master/sdp/sdpOrigin.go#L79)

```go
func (so *SdpOrigin) String() string
```

# transport

```go
import "github.com/KalbiProject/kalbi/transport"
```

## type [ListeningPoint](https://github.com/KalbiProject/kalbi/blob/master/transport/interface.go#L7-L11)

```go
type ListeningPoint interface {
    Read() *message.SipMsg
    Build(string, int)
    Send(string, string, string) error
}
```

### func [NewTransportListenPoint](https://github.com/KalbiProject/kalbi/blob/master/transport/factory.go#L9)

```go
func NewTransportListenPoint(protocol string, host string, port int) ListeningPoint
```

NewTransportListenPoint creates listen

## type [TCPTransport](https://github.com/KalbiProject/kalbi/blob/master/transport/tcp.go#L11-L13)

TCPTransport is a network protocol listening point for the EventDispatcher

```go
type TCPTransport struct {
    // contains filtered or unexported fields
}
```

### func (\*TCPTransport) [Build](https://github.com/KalbiProject/kalbi/blob/master/transport/tcp.go#L32)

```go
func (tt *TCPTransport) Build(host string, port int)
```

### func (\*TCPTransport) [Read](https://github.com/KalbiProject/kalbi/blob/master/transport/tcp.go#L15)

```go
func (tt *TCPTransport) Read() *message.SipMsg
```

## type [UDPTransport](https://github.com/KalbiProject/kalbi/blob/master/transport/udp.go#L12-L15)

UDPTransport is a network protocol listening point for the EventDispatcher

```go
type UDPTransport struct {
    Address    net.UDPAddr
    Connection *net.UDPConn
}
```

### func (\*UDPTransport) [Build](https://github.com/KalbiProject/kalbi/blob/master/transport/udp.go#L28)

```go
func (ut *UDPTransport) Build(host string, port int)
```

### func (\*UDPTransport) [Read](https://github.com/KalbiProject/kalbi/blob/master/transport/udp.go#L18)

```go
func (ut *UDPTransport) Read() *message.SipMsg
```

Read from UDP Socket

### func (\*UDPTransport) [Send](https://github.com/KalbiProject/kalbi/blob/master/transport/udp.go#L42)

```go
func (ut *UDPTransport) Send(host string, port string, msg string) error
```

Generated by [gomarkdoc](https://github.com/princjef/gomarkdoc)
