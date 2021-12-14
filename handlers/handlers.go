package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"

)

type Handlers struct {
	l * log.Logger
}
func NewHello(l *log.Logger) * Handlers{
	return &Handlers{l}
}

func (h *Handlers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")
		d,err:=ioutil.ReadAll(r.Body)
		if err!=nil{
			http.Error(rw, "Oops",http.StatusBadRequest)
			return
		}
	
		fmt.Fprintf(rw,"Hello %s",d)
}
