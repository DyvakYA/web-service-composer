package main

import (
	"./model"
	"./util"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

// main method for starting application
func main() {
	// request handlers
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler);
	router.HandleFunc("/user/firstname/{id}", findUser);
	// request address port
	log.Fatal(http.ListenAndServe(":81", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// bad request if endpoint is "/"
	w.WriteHeader(400)
}

// get user from db
func findUser(w http.ResponseWriter, r *http.Request) {
	// connect to db
	db, err := model.NewDB("root:root@tcp(localhost:3306)/db")
	// catch some errors
	util.Catch(err)
	// define environment
	env := &Env{db: db}
	// check if method request is "GET"
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	// execute method find user
	// and make response
	// in json context-type
	util.RespondwithJSON(w, 200, model.FindUser(env.db, id))
}
