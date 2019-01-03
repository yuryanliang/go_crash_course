package main

import (
	"fmt"
	"net/http"
)

func page1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>page 1<h1>")
}

func page2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "page 2")
}
func main() {
	http.HandleFunc("/index/page1", page1)
	http.HandleFunc("/index/page2", page2)

	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
