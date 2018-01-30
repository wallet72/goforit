package main

import (
	"testing"
	"net/http"
	"fmt"
)

func Testconnect (test *testing.T){

	test.Log("Starting tests...\n\n\n")
	fmt.Printf("Starting testes...\n\n")

	http.HandleFunc("/",helloworld)
	if fail := http.ListenAndServe(":3333",nil); fail != nil {

		test.Errorf("could not open web server",)
		fmt.Printf("Could not open web server")

	}

	fmt.Printf("Web server started...\n\n")

	//defer body.Body.Close()

	body, err := http.Get("127.0.0.1:3333")
	if err != nil {
		test.Errorf("Something wrong with the response from the server")
	}
	fmt.Printf("Response received ok")

	defer body.Body.Close()


}