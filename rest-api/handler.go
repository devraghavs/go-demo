package main

import ("net/http"
"log"
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
		// catch all for any other method then get
	rw.WriteHeader(http.StatusMethodNotAllowed)
	// lp:=GetProducts()
	// //d, err:=json.Marshal(lp)
	// err:= lp.ToJSON(rw)
	// if err!=nil {
	// 	http.Error(rw, "Oops! Unable to marshal json", http.StatusInternalServerError)
	// }
	// // rw.Write(d)
}

func (p*Products) getProducts (rw http.ResponseWriter, r*http.Request){
	lp:=GetProducts()
	err:= lp.ToJSON(rw)
	if err!=nil {
		http.Error(rw, "Oops! Unable to marshal json", http.StatusInternalServerError)
	}
}