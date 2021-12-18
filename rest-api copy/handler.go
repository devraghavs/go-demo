package main

import ("net/http"
"log"
"strconv"
"github.com/gorilla/mux"
"context"
"fmt"
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
	prod:=r.Context().Value(KeyProduct{}).(Product)
	AddProduct(&prod)
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
	// prod:= &Product{}
	// err= prod.FromJSON(r.Body)
	// if err!=nil {
	// 	http.Error(rw, "Oops! Unable to unmarshal json", http.StatusBadRequest)
	// }
	// prod.ID=id
	prod:=r.Context().Value(KeyProduct{}).(Product)
	err=UpdateProduct(id,&prod)
	if err== ErrProductNotFound {
		http.Error(rw, "Oops! Unable to find product", http.StatusNotFound)
		return
	}

	if err!=nil {
		http.Error(rw, "Oops! Unable to update product", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter,r *http.Request){
	prod:= Product{}
	err:= prod.FromJSON(r.Body)
	if err!=nil {
		p.l.Println("Error deserializing product",err)
		http.Error(rw, "Oops! Unable to unmarshal json", http.StatusBadRequest)
	}

	//validate the product
	err= prod.Validate()
	if err!=nil {
		p.l.Println("Error validating product",err)
		http.Error(rw, fmt.Sprintf("Oops! Unable to validate product:%s",err), http.StatusBadRequest)
		return}

	ctx:=context.WithValue(r.Context(),KeyProduct{},prod)
	r = r.WithContext(ctx)
	// ctx:= r.Context().WithValue(KeyProduct{}, prod)
	// req:=r.WithContext(ctx)

	next.ServeHTTP(rw, r)
})
}