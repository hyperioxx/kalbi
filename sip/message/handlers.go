package sip

import "errors"

type methodStorage struct {
	INVITE    func()
	ACK       func()
	BYE       func()
	CANCEL    func()
	REGISTER  func()
	OPTIONS   func()
	PRACK     func()
	SUBSCRIBE func()
	NOTIFY    func()
	PUBLISH   func()
	INFO      func()
	REFER     func()
	MESSAGE   func()
	UPDATE    func()
}

var storage = new(methodStorage)

// HandleFunc sets the function to handle each method
func HandleFunc(method string, function func()) error {
	switch method {
	case "INVITE":
		storage.INVITE = function
	case "ACK":
		storage.ACK = function
	case "BYE":
		storage.BYE = function
	case "CANCEL":
		storage.CANCEL = function
	case "REGISTER":
		storage.REGISTER = function
	case "OPTIONS":
		storage.OPTIONS = function
	case "PRACK":
		storage.PRACK = function
	case "SUBSCRIBE":
		storage.SUBSCRIBE = function
	case "NOTIFY":
		storage.NOTIFY = function
	case "PUBLISH":
		storage.PUBLISH = function
	case "INFO":
		storage.INFO = function
	case "REFER":
		storage.REFER = function
	case "MESSAGE":
		storage.MESSAGE = function
	case "UPDATE":
		storage.UPDATE = function
	default:
		return errors.New("Unkown Method Selected")
	}
	return nil
}

// GetMethodFunction returns function
func GetMethodFunction(method string) func() {

	switch method {
	case "INVITE":
		return storage.INVITE
	case "ACK":
		return storage.ACK
	case "BYE":
		return storage.BYE
	case "CANCEL":
		return storage.CANCEL
	case "REGISTER":
		return storage.REGISTER
	case "OPTIONS":
		return storage.OPTIONS
	case "PRACK":
		return storage.PRACK
	case "SUBSCRIBE":
		return storage.SUBSCRIBE
	case "NOTIFY":
		return storage.NOTIFY
	case "PUBLISH":
		return storage.PUBLISH
	case "INFO":
		return storage.INFO
	case "REFER":
		return storage.REFER
	case "MESSAGE":
		return storage.MESSAGE
	case "UPDATE":
		return storage.UPDATE
	default:
		return storage.CANCEL
	}

}
