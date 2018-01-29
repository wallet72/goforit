package main

import (
	"testing"
	"net/http"
)

func Testconnect (t *testing.T){

	http.HandleFunc("/",helloworld)
	if err := http.ListenAndServe(":3333",nil); err != nil {

		t.Errorf("could not open web server",)

	}

}