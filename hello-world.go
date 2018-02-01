package main

import "net/http"


func webserver(web http.ResponseWriter, req *http.Request){
	web.WriteHeader(200)
	web.Write([]byte("Hello World"))

}

func main(){
	http.HandleFunc("/",webserver)
	http.ListenAndServe(":3333",nil)

}