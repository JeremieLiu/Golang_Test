package main

import (
	"net/http"
	"fmt"
	"log"
)

func main(){
	http.HandleFunc("/testWord",initFun)
	Path := http.FileServer(http.Dir("C://0.work//golang_workspace//"))
	http.Handle("/static/",http.StripPrefix("/static/",Path))
	if Path == nil {
		log.Fatal(Path)
	}


	err:=http.ListenAndServe(":7085",nil)
	if err == nil {
		log.Fatal(err)
	}

}

func initFun(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"test for display")
}
