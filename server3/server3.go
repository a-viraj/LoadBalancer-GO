package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8003"

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Server started at %s",port)
	})
	log.Fatal(http.ListenAndServe(":"+port,nil))
}