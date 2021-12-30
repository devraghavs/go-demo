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
	
	Order(string,map[string][]float64) (string, error)
}

// orderService is a concrete implementation of StringService
type orderService struct{}

func (orderService) Order(l string ,arr map[string][]float64) (string, error) {
    total := 0.0
    if len(l)!=4 || !(strings.ContainsAny(l,"A | B | C") && strings.ContainsAny(l,"D | E") && strings.ContainsAny(l,"F | G | H") && strings.ContainsAny(l,"I | J")){
        return "",ErrEmpty
    }else{
        for i:=range l{
            key:=string(l[i])
            if arr[key][1]==0{
				return "",ErrEmptyPart
                break
            }else{
                    total+=arr[key][0]
                    arr[key][1]=arr[key][1] - 1
            }
        }
        x := fmt.Sprintf("%v", total)
        return x,nil
	}
}



// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("Invalid Input")
var ErrEmptyPart = errors.New("Error : Cannot create robot")

// For each method, we define request and response structs
type orderRequest struct {
	S string `json:"s"`
}

type ordercaseResponse struct {
	V   string `json:"total,omitempty"`
	
	Err string `json:"Error:,omitempty"` // errors don't define JSON marshaling
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeOrderEndpoint(svc StringService) endpoint.Endpoint {
	arr:= map[string][]float64{
        "A":{10.28,9},
        "B":{24.07,7},
        "C":{13.30,0},
        "D":{28.94,1},
        "E":{12.39,3},
        "F":{30.77,2},
        "G":{55.13,15},
        "H":{50.00,7},
        "I":{90.12,92},
        "J":{82.31,15},
    }
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(orderRequest)
		v, err := svc.Order(req.S,arr)
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