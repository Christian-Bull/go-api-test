package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InputList is a list of input types
type InputList struct {
	Input []Input
}

// Input is the data struct used for a request
type Input struct {
	Data string `json:"data"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey there ( ͡° ͜ʖ ͡°)")
	fmt.Println("Endpoint hit: home")
}

func reverseString(w http.ResponseWriter, r *http.Request) {
	// unmarshel the request body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var inputs []Input
	err = json.Unmarshal(b, &inputs)

	// reverse the string
	for i := 0; i < len(inputs); i++ {
		inputs[i].Data = reverse(inputs[i].Data)
	}

	// Returns in the same format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inputs)

}
