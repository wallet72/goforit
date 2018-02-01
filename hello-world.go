package main

import "net/http"

// This is a really simple way to start a brain dead web server
// Thanks to the net/http module.
func webserver(web http.ResponseWriter, req *http.Request){
	web.WriteHeader(http.StatusOK)
	web.Write([]byte("Hello World"))

}

// So listen and respond with the message...
func main(){
	http.HandleFunc("/",webserver)
	http.ListenAndServe(":3333",nil)

}