package sip

import (
	    "fmt"
	    "regexp"
		"strings"
		//"bufio"
	   )

type SipMessage struct {
	Raw string 
	Headers map[string][]SIPHeader
}


func (m *SipMessage) Parse(message string){
	pattern := regexp.MustCompile("([a-zA-Z-]+):(.*)")
	
	for _, line := range strings.Split(message, "\n") {
		fmt.Println(pattern.FindString(line))
	}

}