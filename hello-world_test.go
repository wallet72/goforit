package main

import (
	"testing"
	"net/http"
	"fmt"
)

func Testconnect (t *testing.T){

	fmt.Printf("Starting tests...\n\n\n")

	http.HandleFunc("/",helloworld)
	if fail := http.ListenAndServe(":3333",nil); fail != nil {

		t.Errorf("could not open web server",)

	}

	body, err := http.Get("127.0.0.1:3333")
	if err != nil {
		t.Errorf("Something wrong with the response from the server")
	}

	defer body.Body.Close()


}