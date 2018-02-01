package main

import (
	"testing"
	"net/http"
	"fmt"
	"log"
	"net/http/httptest"
)

func responsetest (rtest *testing.T){

	rtarget,rerror := http.NewRequest("GET","/",nil)
	if rerror != nil {
		fmt.Printf("Prep work failed: %s,%s", rtarget,rerror)
		rtest.Fatal(rerror)
	}

	incoming := httptest.NewRecorder()
	processor := http.HandlerFunc(webserver)
	processor.ServeHTTP(incoming,rtarget)

	status := incoming.Code;

	if status != http.StatusOK {
		rtest.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if incoming.Body.String() != "Hello World" {
		rtest.Errorf("Message body was not right: %v", incoming.Body.String())
	}
}
