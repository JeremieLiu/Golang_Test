package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	//调用post方法
	httpPostForm()
}

func httpPostForm() {

	resp, err := http.PostForm("http://localhost:7081/dis1",
		url.Values{"first name": {"jeremie"}, "last name": {"Liu"}})
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
