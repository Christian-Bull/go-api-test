package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/reverse", reverseString)

	log.Fatal(http.ListenAndServe(":5050", myRouter))
}
