package transaction

import (
	"github.com/KalbiProject/Kalbi/sip/message"
	"github.com/KalbiProject/Kalbi/transport"
)

/*
   Author - Aaron Parfitt
   Date - 11th October 2020


   RFC3261 - SIP: Session Initiation Protocol
   https://tools.ietf.org/html/rfc3261#section-17

   Transactions

   SIP is a transactional protocol: interactions between components take
   place in a series of independent message exchanges.  Specifically, a
   SIP transaction consists of a single request and any responses to
   that request, which include zero or more provisional responses and
   one or more final responses.  In the case of a transaction where the
   request was an INVITE (known as an INVITE transaction), the
   transaction also includes the ACK only if the final response was not
   a 2xx response.  If the response was a 2xx, the ACK is not considered
   part of the transaction.

      The reason for this separation is rooted in the importance of
      delivering all 200 (OK) responses to an INVITE to the UAC.  To
      deliver them all to the UAC, the UAS alone takes responsibility

      for retransmitting them (see Section 13.3.1.4), and the UAC alone
      takes responsibility for acknowledging them with ACK (see Section
      13.2.2.4).  Since this ACK is retransmitted only by the UAC, it is
      effectively considered its own transaction.

   Transactions have a client side and a server side.  The client side
   is known as a client transaction and the server side as a server
   transaction.  The client transaction sends the request, and the
   server transaction sends the response.  The client and server
   transactions are logical functions that are embedded in any number of
   elements.  Specifically, they exist within user agents and stateful
   proxy servers.  Consider the example in Section 4.  In this example,
   the UAC executes the client transaction, and its outbound proxy
   executes the server transaction.  The outbound proxy also executes a
   client transaction, which sends the request to a server transaction
   in the inbound proxy.  That proxy also executes a client transaction,
   which in turn sends the request to a server transaction in the UAS.
   This is shown in Figure 4.

   +---------+        +---------+        +---------+        +---------+
   |      +-+|Request |+-+   +-+|Request |+-+   +-+|Request |+-+      |
   |      |C||------->||S|   |C||------->||S|   |C||------->||S|      |
   |      |l||        ||e|   |l||        ||e|   |l||        ||e|      |
   |      |i||        ||r|   |i||        ||r|   |i||        ||r|      |
   |      |e||        ||v|   |e||        ||v|   |e||        ||v|      |
   |      |n||        ||e|   |n||        ||e|   |n||        ||e|      |
   |      |t||        ||r|   |t||        ||r|   |t||        ||r|      |
   |      | ||        || |   | ||        || |   | ||        || |      |
   |      |T||        ||T|   |T||        ||T|   |T||        ||T|      |
   |      |r||        ||r|   |r||        ||r|   |r||        ||r|      |
   |      |a||        ||a|   |a||        ||a|   |a||        ||a|      |
   |      |n||        ||n|   |n||        ||n|   |n||        ||n|      |
   |      |s||Response||s|   |s||Response||s|   |s||Response||s|      |
   |      +-+|<-------|+-+   +-+|<-------|+-+   +-+|<-------|+-+      |
   +---------+        +---------+        +---------+        +---------+
      UAC               Outbound           Inbound              UAS
                        Proxy               Proxy

                  Figure 4: Transaction relationships

   A stateless proxy does not contain a client or server transaction.
   The transaction exists between the UA or stateful proxy on one side,
   and the UA or stateful proxy on the other side.  As far as SIP
   transactions are concerned, stateless proxies are effectively
   transparent.  The purpose of the client transaction is to receive a
   request from the element in which the client is embedded (call this
   element the "Transaction User" or TU; it can be a UA or a stateful
   proxy), and reliably deliver the request to a server transaction.

   The client transaction is also responsible for receiving responses
   and delivering them to the TU, filtering out any response
   retransmissions or disallowed responses (such as a response to ACK).
   Additionally, in the case of an INVITE request, the client
   transaction is responsible for generating the ACK request for any
   final response accepting a 2xx response.

   Similarly, the purpose of the server transaction is to receive
   requests from the transport layer and deliver them to the TU.  The
   server transaction filters any request retransmissions from the
   network.  The server transaction accepts responses from the TU and
   delivers them to the transport layer for transmission over the
   network.  In the case of an INVITE transaction, it absorbs the ACK
   request for any final response excepting a 2xx response.

   The 2xx response and its ACK receive special treatment.  This
   response is retransmitted only by a UAS, and its ACK generated only
   by the UAC.  This end-to-end treatment is needed so that a caller
   knows the entire set of users that have accepted the call.  Because
   of this special handling, retransmissions of the 2xx response are
   handled by the UA core, not the transaction layer.  Similarly,
   generation of the ACK for the 2xx is handled by the UA core.  Each
   proxy along the path merely forwards each 2xx response to INVITE and
   its corresponding ACK.

*/

//Transaction
type Transaction interface {
	GetBranchID() string
	GetOrigin() *message.SipMsg
	SetListeningPoint(transport.ListeningPoint)
	Send(*message.SipMsg, string, string)
	Receive(*message.SipMsg)
	GetLastMessage() *message.SipMsg
	GetServerTransactionID() string
   SetLastMessage(*message.SipMsg)
   GetListeningPoint() transport.ListeningPoint
}
