package test

import (
	"net/http"
	"io"
)

func helloHandler(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "hello world \n")
}



func main(){
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":12345", nil)
}