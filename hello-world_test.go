package main

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
)

// Define the testing routine to use the brain dead listener
// We are testing that the status is OK and that the
// message body is what we expected
func TestHelloworld (rtest *testing.T){

	//Get it...
	rtarget,rerror := http.NewRequest("GET","/",nil)
	if rerror != nil {
		fmt.Printf("Prep work failed: %s,%s", rtarget,rerror)
		rtest.Fatal(rerror)
	}

	incoming := httptest.NewRecorder()
	processor := http.HandlerFunc(webserver)
	processor.ServeHTTP(incoming,rtarget)

	status := incoming.Code;

	//test the status
	if status != http.StatusOK {
		rtest.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//test the message
	if incoming.Body.String() != "Hello World" {
		rtest.Errorf("Message body was not right: %v", incoming.Body.String())
	}
}
