package main

import (
	"fmt"
	"kalbi"
)

type Handler struct {
}

func (h *Handler) ServeSIP(w kalbi.ResponseWriter, r *kalbi.Request) {
    fmt.Println(r)
}

func main() {
	serv := kalbi.NewUDPServer("0.0.0.0:5060", &Handler{})

	serv.ListenAndServe()
}
