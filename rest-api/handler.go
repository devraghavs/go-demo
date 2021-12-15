package main

import ("net/http"
"log"
"encoding/json"
)

type Products struct {
l* log.Logger

}

func NewProducts(l* log.Logger) *Products {
	return &Products{l}
}

func(p*Products) ServeHTTP(rw http.ResponseWriter, r*http.Request) {
	lp:=GetProducts()
	json.Marshal(lp)
	d, err:=json.Marshal(lp)
	if err!=nil {
		http.Error(rw, "Oops! Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}