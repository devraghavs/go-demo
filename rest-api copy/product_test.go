package main

import "testing"

func TestChecksValidation(t *testing.T){
	p:= &Product{
		Name:"Hy",
		Price: 1.00,
		SKU: "aba-sss-srr"
	}

	err:=p.Validate()

	if err!=nil {
		t.Fatal(err)
	}
}