package main

import (
	"testing"
	"net/http"
	"fmt"
	//"./hello-world"
	"log"
)

func Testconnect (test *testing.T) {

	test.Log("Starting tests...\n\n\n")
	fmt.Printf("Starting testes...\n\n")

	http.HandleFunc("/", webserver)
	if fail := http.ListenAndServe(":3333", nil); fail != nil {

		test.Errorf("could not open web server", )
		fmt.Printf("Could not open web server", )

	}

	fmt.Printf("Web server started...\n\n")

	webresponse, badness := http.Get("127.0.0.1:3333")

	if badness != nil {

		fmt.Printf("Error response %s", badness)
		log.Fatal(badness)

	}

	/*if webresponse != "hello world"{

		log.Fatal("invalid response received: $S", webresponse)
	}*/
	fmt.Printf("Response received: %s\n\n", webresponse)
	fmt.Printf("testing passed ok",)

}