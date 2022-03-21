package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

type salePriceresp struct {
	Price float64 `json:"salePrice"`
}

func getProductPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sku := vars["sku"]
	urls := "https://api.bestbuy.com/v1/products/" + sku + ".json?&show=salePrice&apiKey=WsZYk44qBGD8qima1OTZ59Gg"
	resp, err := http.Get(urls)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Fatalln(err)
	}
	var result salePriceresp
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result.Price)
	w.Write(body)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products/{sku}", getProductPrice).Methods("GET")
	http.ListenAndServe(":8080", r)
}
