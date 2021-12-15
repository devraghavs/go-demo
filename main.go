// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"log"
// 	"io/ioutil"
// )

// func main() {
// 	http.HandleFunc("/",func(rw http.ResponseWriter,r*http.Request){
// 		log.Println("Hello World")
// 		d,err:=ioutil.ReadAll(r.Body)
// 		if err!=nil{
// 			http.Error(rw, "Oops",http.StatusBadRequest)
// 			return
// 		}
// 		fmt.Fprintf(rw,"Hello %s",d)
// 	})

// 	http.HandleFunc("/goodbye",func(rw http.ResponseWriter, r * http.Request){
// 		log.Println("Goodbye World")
// 	})
// 	http.ListenAndServe(":8080",nil)
// }
package main

import (
	"GoLang/handlers"
	"net/http"
	"log"
)

func main() {
	l:=log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHell(l)
	
	sm:= NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":8080",sm)
}