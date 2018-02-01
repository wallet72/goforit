package main

import (
	"testing"
	"net/http"
	"fmt"
	"./hello-world"
	"log"
)

func Testconnect (test *testing.T) {

	test.Log("Starting tests...\n\n\n")
	fmt.Printf("Starting testes...\n\n")

	http.HandleFunc("/", helloworld)
	if fail := http.ListenAndServe(":3333", nil); fail != nil {

		test.Errorf("could not open web server", )
		fmt.Printf("Could not open web server", )

	}
}

func webtest(handle *testing.T){

	fmt.Printf("Web server started...\n\n")

	webresponse, error := http.Get("127.0.0.1:3333")

	if error != nil {

		log.Fatal("error: %s", error)

	}

	if webresponse != "hello world"{

		log.Fatal("invalid response received: %S", webresponse)
	}

	fmt.Printf("testing passed ok",)

}