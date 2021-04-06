package main

import (
	"challenge-golang-stone/src/config"
	"challenge-golang-stone/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.Load()
}

func main() {
	router := router.Generate()
	port := config.Port
	fmt.Printf("Api is running on port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
