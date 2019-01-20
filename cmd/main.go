package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/console", consoleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func consoleHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != "POST" {
		w.WriteHeader(500)
	}

	raw, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// handle
		log.Println(err)
		// use 500 Internal Server Error
		w.WriteHeader(500)
	}

	// "{key: value, key_1:value_1}"
	creds := Credentials{}

	err = json.Unmarshal(raw, &creds)
	if err != nil {
		// we fail to unmarshal it
		log.Println(err)
		w.WriteHeader(500)
	}

	fmt.Println(creds)
}

type Credentials struct {
	Username string `json:"username"`
	Secret   string `json:"password"`
}
