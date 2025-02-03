//req headers form http req
//ttype Header map[string][] string
//using header.vlaues 
//func (h header)Values(key string)[]string
//func(h header) Get(key string)string
//direct accces 
//map[caccept]
//
package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8080", nil)
}
func handleRequest(w http.ResponseWriter, r *http.Request) {

	fooVAlues := r.Header.Values("foo")
	fmt.Printf("r.Header.Values(\"foo\"):: %s\n\n", fooVAlues)

	contentType := r.Header.Get("content-type")
	fmt.Printf("r.Header.Get(\"content-type\"):: %s\n\n", contentType)

	fooFirstValue := r.Header.Get("foo")
	fmt.Printf("r.Header.Get(\"foo\"):: %s\n\n", fooFirstValue)

	headers := r.Header
	fmt.Printf("r.Headers:: %s\n\n", headers)
	fmt.Printf("r.Headers[\"Content-Type\"]:: %s\n\n", headers["Content-Type"])
	fmt.Printf("r.Headers[\"Foo\"]:: %s", headers["Foo"])
}