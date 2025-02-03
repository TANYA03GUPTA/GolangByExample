package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	input_url := "http://localhost:8080/products?filters=color&filters=price&order=asc"
	u, err := url.Parse(input_url)
	if err != nil {
		log.Fatal(err)
	}

	queryParams := u.Query()

	fmt.Println(queryParams)

	fmt.Println(queryParams["filters"])

	fmt.Println(queryParams["order"])

}