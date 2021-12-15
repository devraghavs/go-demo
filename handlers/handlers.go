package main

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
// Second Handler i.e. goodbye handler

type Goodbye struct {
	l * log.Logger
}
func NewGoodbye(l *log.Logger) * Goodbye{
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

rw.Write([]byte("Goodbye World"))}


