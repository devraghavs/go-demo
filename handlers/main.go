package main

import (
	 "GoLang/handlers"
	"net/http"
	"log"
)

func main() {
	l:=log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	
	sm:= NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":8080",sm)
}