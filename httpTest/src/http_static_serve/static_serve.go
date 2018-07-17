package main

import (
	//"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/",display_1)
	//serveMux.HandleFunc("/static/",httpServe)

	httpServe()

	server := http.Server{
		Addr:		string(":7100"),
		Handler:	serveMux,
		ReadTimeout:	3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func display_1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

//func httpServe(w http.ResponseWriter,r *http.Request){
//	wd, err := os.Getwd()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	http.StripPrefix("/static_server", http.FileServer(http.Dir(wd))).ServeHTTP(w,r)
//}

func httpServe(){
	http.Handle("/test",http.FileServer(http.Dir("file")))

	err := http.ListenAndServe(":7100",nil)
	if err != nil {
		log.Fatal(err)
	}
}


