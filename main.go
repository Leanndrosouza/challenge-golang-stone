package main

import (
	"challenge-golang-stone/src/config"
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func init() {
	config.Load()
}

func main() {
	http.HandleFunc("/", helloWorld)
	port := config.Port
	fmt.Printf("Api is running on port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
