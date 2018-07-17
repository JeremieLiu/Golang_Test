package main

import (
	"net/http"
	"log"
	"fmt"
)

func main(){
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/static/",fileServer)

	err := http.ListenAndServe("localhost:7081",nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"test for static serve")
}

func fileServer(w http.ResponseWriter , r *http.Request){
	http.ServeFile(w, r, "static_serve.go")
	//r.URL.Path[1:]
}