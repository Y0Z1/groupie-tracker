package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/styles/",http.StripPrefix("/styles/",http.FileServer(http.Dir("styles"))))
	http.HandleFunc("/",homepage)
	http.HandleFunc("/artist",artistpage)
	http.HandleFunc("/error-404",errorpage404)
	http.HandleFunc("/error-400",errorpage400)
	fmt.Println("localhost:8080 Server is running...")
	http.ListenAndServe(":8080",nil)
}