package main
 
import ("time"
"encoding/json"
"io"
"fmt"
)

//Product defines the structure of the product

type Product struct {
	ID int        `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}
type Productss []*Product

func (p*Productss) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p*Product) FromJSON(r io.Reader) error {
	e:=json.NewDecoder(r)
	 return e.Decode(p)
}

func GetProducts() Productss {
	return productList
}

func AddProduct(p *Product){
	p.ID=getNextID()
	productList=append(productList,p)
}
var ErrProductNotFound = fmt.Errorf("Product Not Found")

func UpdateProduct(id int, p *Product) error{
	_,pos,err:=findProduct(id)
	if err!=nil {
		return err
	
	}

	p.ID=id
	productList[pos]=p
	return nil
}

func findProduct(id int) (*Product, int, error){
	for i,p:=range productList{
		if p.ID==id{
			return p,i,nil}

	}
	return nil,-1,ErrProductNotFound

}

func getNextID() int {
	lp:=productList[len(productList)-1]
	return lp.ID + 1
	
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 3.00,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}