package main

import (
	"encoding/json"
	"net/http"
)

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/posts", handleGetList)
	http.HandleFunc("/posts/", handleRequest)

	server.ListenAndServe()
}


func handleGetList(w http.ResponseWriter, r *http.Request){
	var err error

	posts, err := getPosts(100)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	output, err := json.MarshalIndent(&posts, "", "\t")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().set("Access-Control-Allow-Headers", "*")
	w.Header().set("Access-Control-Allow-Origin", "*")
	w.Header().set("Access-Control-Allow-Methods", "GET")
	w.Write(output)
}


func handleRequest(w http.ResponseWriter, r *http.Request){
	var err error

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = hadlePost(w, r)
	case "PUT":
		err = hadlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}