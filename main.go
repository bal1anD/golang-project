package main

import (
	"fmt"
	"gawkbox-assignment/myrouter"
	"net/http"

)

func main() {
	fmt.Println("Booting the server...")
	http.HandleFunc("/", myrouter.Route)

	// Run your server
	http.ListenAndServe(":8080", nil)
}
