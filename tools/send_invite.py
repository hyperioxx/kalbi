import socket

def send_sip_invite(ip, port):
    # Prepare the sender and receiver details
    sender_uri = "sip:sender@192.168.1.100"
    receiver_uri = "sip:user@{}".format(ip)
    
    # Create the SIP INVITE message
    invite_msg = f"""INVITE {receiver_uri} SIP/2.0
Via: SIP/2.0/UDP 192.168.1.100:5060;branch=z9hG4bKnashds7
Max-Forwards: 70
To: <{receiver_uri}>
From: "Sender" <{sender_uri}>;tag=7743
Call-ID: 1234567890@192.168.1.100
CSeq: 1 INVITE
Contact: <{sender_uri}>
Content-Type: application/sdp
Content-Length: 0

"""

    # Create a socket
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    # Send the INVITE message
    sock.sendto(invite_msg.encode(), (ip, port))
    print(f"SIP INVITE sent to {ip}:{port}")

    # Close the socket
    sock.close()

# Example usage:
send_sip_invite("localhost", 5060)
