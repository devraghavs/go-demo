package main

import(
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Request received")
		fmt.Println("Method:", r.Method)
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe("localhost:3000", mux)

}