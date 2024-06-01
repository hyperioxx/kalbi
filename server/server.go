package server

type Handler interface {
	Serve(ResponseWriter, *Request)
}

type Server interface {
	ListenAndServe()
}