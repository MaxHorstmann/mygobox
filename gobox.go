package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("listing...")
	err := http.ListenAndServe(":1234", nil)

}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "test")
}