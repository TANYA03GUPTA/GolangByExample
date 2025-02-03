//jhow to fetch params form a query ?
//1. way is to use Query()
//it stores key value
//2. is to get value form keyname
//3.   r.ParseForm() use this to get map pof all fileters
package main

import (
    "net/http"
)

func main() {
    getProductsHandler := http.HandlerFunc(getProducts)
    http.Handle("/products", getProductsHandler)
    http.ListenAndServe(":8080", nil)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    filters := query.Get("filters") //filters="color"
    w.WriteHeader(200)
    w.Write([]byte(filters))
	r.ParseForm()
	filters, present := r.Form["filters"]

	filters := r.FormValue("filters") 

	query = r.URL.Query()
    filters, present := query["filters"] //filters=["color", "price", "brand"]
    if !present || len(filters) == 0 {
        fmt.Println("filters not present")
    }
}