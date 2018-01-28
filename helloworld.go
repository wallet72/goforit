package main

import (
	"net/http"
	"fmt"
)

func helloworld(web http.ResponseWriter, req *http.Request){
	fmt.Println(web,"Hello", "world")
}

func main(){
	http.HandleFunc("/",helloworld)
	http.ListenAndServe("localhost:80",nil)

}