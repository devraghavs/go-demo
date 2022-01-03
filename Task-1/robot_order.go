package main

import (
	"fmt"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// StringService provides operations on strings.
type StringService interface {
	Order(string) (string, error)
}

// orderService is a concrete implementation of StringService
type orderService struct{}

func (orderService) Order(l string) (string, error) {
	arr:= [10][3]interface{}{
        {"A",10.28,9},
        {"B",24.07,7},
        {"C",13.30,0},
        {"D",28.94,1},
        {"E",12.39,3},
        {"F",30.77,2},
        {"G",55.13,15},
        {"H",50.00,7},
        {"I",90.12,92},
        {"J",82.31,15},
    }
	total := 0.0
    res1 :=0
    flag :=0;
    var flag2 int;
	if l == "" {
			return "", ErrEmpty
		}
    if len(l)!=4 || !(strings.ContainsAny(l,"a | b | c") && strings.ContainsAny(l,"d | e") && strings.ContainsAny(l,"f | g | h") && strings.ContainsAny(l,"i | j")){
        return ErrInput,nil;
    }else{
         for i:=range l{
             for j:=range arr{
             res:=strings.Index(arr[j][0].(string),strings.ToUpper(string(l[i])))
             if res==0{
                 res1=j
                 if arr[j][2].(int)>0{
                    fmt.Println("res1",res1,arr[j][1]) 
                    total+=arr[res1][1].(float64)
                    // arr[j][2].(int)=(arr[j][2].(int))-1
                    break
                 }else{
                      flag2=1
                     break
                 }}}
            if flag2<1{
				continue
            }else{
                break}}
     if flag2<1 && flag!=1{
		x := fmt.Sprintf("%v", total)
         return x,nil
    }else{
        return ErrEmptyPart,nil}}
}
// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
var ErrEmptyPart = "Error : Cannot create robot"
var ErrInput="Invalid Input:Please enter all types of body part"

// For each method, we define request and response structs
type orderRequest struct {
	S string `json:"s"`
}

type ordercaseResponse struct {
	V   string `json:"total"`
	
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeOrderEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(orderRequest)
		v, err := svc.Order(req.S)
		if err != nil {
			return ordercaseResponse{v, err.Error()}, nil
		}
		return ordercaseResponse{v, ""}, nil
	}
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := orderService{}
	orderHandler := httptransport.NewServer(
		makeOrderEndpoint(svc),
		decodeOrderRequest,
		encodeResponse,
	)
	http.Handle("/order", orderHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request orderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	
	return json.NewEncoder(w).Encode(response)
}