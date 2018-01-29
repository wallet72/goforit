package main

import (
	"net/http"
	//"fmt"
	//"strings"
)

func helloworld(web http.ResponseWriter, req *http.Request){

	//message = "Hello World"

	web.Write([]byte("message"))


	//fmt.Println(web,"Hello", "world")
}

func main(){
	http.HandleFunc("/",helloworld)
	http.ListenAndServe(":3030",nil)

}