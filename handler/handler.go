package handler

import (
	"net/http"
	"io"
	"github.com/globalsign/mgo"
	//"io/ioutil"
	"log"
	"encoding/json"
	//"io/ioutil"
	"cloudnative/data"
	"fmt"
	"github.com/globalsign/mgo/bson"
)

func MainHandler(w http.ResponseWriter, r *http.Request){
	//fmt.Println(w, "welcome to first page")
	//w.Header().Set()
	io.WriteString(w, "Welcome to home page\n")
}

func MakeUserHandler(w http.ResponseWriter, r *http.Request){
	//username := r.Body.username
	//c.Insert(&user{username:,password:,nickname:})
}

func AddUserHandler(s *mgo.Session)func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		c := session.DB("golang").C("program")
		// to parse the body from http request
		decoder := json.NewDecoder(r.Body)
		var newuser data.User
		//body, _ := ioutil.ReadAll(r.Body)
		err := decoder.Decode(&newuser)
		fmt.Println("newuser name is ", newuser.Username)
		if err !=nil {
			log.Fatal("erro parsing json")
		}
		fmt.Printf("stop position")
		//c.Insert(&data.User{user["username"],user["password"],user["nickname"]})
		c.Insert(&newuser)
		io.WriteString(w, "add user done \n")
	}
}

func DeleteUserHandler(s *mgo.Session)func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
	    //session := s.Copy()
	    //c := session.DB("golang").C("program")
	    ////json_body := json.NewDecoder()
	    //c.Drop
	}
}

func GetUserHandler(s *mgo.Session)func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		c := s.DB("golang").C("program")
		var userlist []data.User
		err := c.Find(bson.M{}).All(&userlist)
		if err != nil{
			log.Fatal("err is true")
		}
		//io.WriteString(w, json.Marshal(userlist))
		response_data,err := json.Marshal(userlist)
		io.WriteString(w, string(response_data))
	}
}