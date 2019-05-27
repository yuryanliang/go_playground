package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	HandleRequest()
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":1111", nil))
}

func homePage(responseWriter http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(responseWriter, "Welcome to the HomePage!")
	_, _ = io.WriteString(responseWriter, "{\"status\":\"good\"}")

	fmt.Println("Endpoint Hit: homePage")
}
