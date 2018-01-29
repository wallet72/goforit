package main

import "net/http"


func helloworld(web http.ResponseWriter, req *http.Request){

	web.Write([]byte("Hello World"))

}

func main(){
	http.HandleFunc("/",helloworld)
	http.ListenAndServe(":3333",nil)

}