package main

import (
	"log"
	"net/http"
)

func main(){

	//.ServeHTTP(w http.ResponseWriter, r *http.Request{})
	http.Handle("/",http.FileServer(http.Dir("tets.go")))
	err := http.ListenAndServe(":7085",nil)
	if err != nil {
		log.Fatal(err)
	}

}

//func fileServer(w http.ResponseWriter , r *http.Response){
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//	http.StripPrefix("/static_server", http.FileServer(http.Dir(wd))).ServeHTTP(w,r)
//}