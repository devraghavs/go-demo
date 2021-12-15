package main

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"time"
	"context"
)

func main() {
	l:=log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := NewProducts(l)	
	sm:= http.NewServeMux()
	sm.Handle("/", ph)

	s:= &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second, 

	}
	go func() {
	err:=s.ListenAndServe()
	if err!=nil{
		l.Fatal(err)
	}
}()
sigChan:= make(chan os.Signal)
signal.Notify(sigChan, os.Interrupt)
signal.Notify(sigChan,os.Kill)

sig:=<-sigChan
l.Println("Recieved shutdown signal, exiting",sig)

	tc,_:=context.WithTimeout(context.Background(),30*time.Second)
	s.Shutdown(tc)
	// http.ListenAndServe(":8080",sm)
}