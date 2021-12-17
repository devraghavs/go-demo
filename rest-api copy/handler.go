package main

import ("net/http"
"log"
"strconv"
"github.com/gorilla/mux"
)

type Products struct {
l* log.Logger

}

func NewProducts(l* log.Logger) *Products {
	return &Products{l}
}


func (p*Products) GetProducts (rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle getProducts")
	lp:=GetProducts()
	err:= lp.ToJSON(rw)
	if err!=nil {
		http.Error(rw, "Oops! Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p* Products) AddProduct (rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle addProduct")
	prod:= &Product{}
	err:= prod.FromJSON(r.Body)
	if err!=nil {
		http.Error(rw, "Oops! Unable to unmarshal json", http.StatusBadRequest)
	}
	AddProduct(prod)
	//p.l.Printf("Prod: %#v", prod)

}

func (p* Products) UpdateProducts (rw http.ResponseWriter, r*http.Request){
	vars:= mux.Vars(r)
	id,err:=strconv.Atoi(vars["id"])
	if err!=nil{
		http.Error(rw, "Oops! Unable to convert id to int", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle updateProducts",id)
	prod:= &Product{}
	err= prod.FromJSON(r.Body)
	if err!=nil {
		http.Error(rw, "Oops! Unable to unmarshal json", http.StatusBadRequest)
	}
	prod.ID=id
	err=UpdateProduct(id,prod)
	if err== ErrProductNotFound {
		http.Error(rw, "Oops! Unable to find product", http.StatusNotFound)
		return
	}

	if err!=nil {
		http.Error(rw, "Oops! Unable to update product", http.StatusInternalServerError)
		return
	}



}