package main

import (
	"fmt"
	"hello-world/pkg/handlers"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/div", Divide)
	fmt.Printf("Starting app : %s%s", "http://localhost", port)
	_ = http.ListenAndServe(port, nil)
}
