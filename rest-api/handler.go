package main

import ("net/http"
"log"
"regexp"
"strconv"
)

type Products struct {
l* log.Logger

}

func NewProducts(l* log.Logger) *Products {
	return &Products{l}
}

func(p*Products) ServeHTTP(rw http.ResponseWriter, r*http.Request) {

	if r.Method==http.MethodGet{
		p.getProducts(rw,r)
		return
	}
	if r.Method==http.MethodPost{
		p.addProduct(rw,r)
		return
	}

	if r.Method == http.MethodPut {
		reg:=regexp.MustCompile(`/([0-9]+)`)
		g:=reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g)!=1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return }
		if len(g[0])!=2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString:=g[0][1]
		id,err:=strconv.Atoi(idString)
		if err!=nil {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return 
		}
		// p.l.Println(id)
		p.updateProducts(rw,r,id)
		return	
	}

		// catch all for any other method then get
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) getProducts (rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle getProducts")
	lp:=GetProducts()
	err:= lp.ToJSON(rw)
	if err!=nil {
		http.Error(rw, "Oops! Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p* Products) addProduct (rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle addProduct")
	prod:= &Product{}
	err:= prod.FromJSON(r.Body)
	if err!=nil {
		http.Error(rw, "Oops! Unable to unmarshal json", http.StatusBadRequest)
	}
	AddProduct(prod)
	//p.l.Printf("Prod: %#v", prod)

}

func (p* Products) updateProducts (rw http.ResponseWriter, r*http.Request, id int){
	p.l.Println("Handle updateProducts")
	prod:= &Product{}
	err:= prod.FromJSON(r.Body)
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