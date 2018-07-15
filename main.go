package main

import (
	"net/http"
	"log"
	"github.com/globalsign/mgo"
	handler "cloudnative/handler"
)



func main(){
	session, err := mgo.Dial("localhost:27017")
	if err!=nil {
		log.Fatal("mongo conection error")
	}
	defer session.Close()
	//c := session.DB("golang").C("program")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.MainHandler)
	mux.HandleFunc("/adduser", handler.AddUserHandler(session))
	mux.HandleFunc("/deluser", handler.DeleteUserHandler(session))
	mux.HandleFunc("/getuser", handler.GetUserHandler(session))
	rh := http.RedirectHandler("http://www.163.com", 307)
	mux.Handle("/redict", rh)
	log.Println("Listening")
	log.Fatal(http.ListenAndServe(":8888", mux))
}