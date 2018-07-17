package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//set router
	http.HandleFunc("/dis1", display_1)
	http.HandleFunc("/dis2", display_2)

	//set command
	err := http.ListenAndServe(":7081", nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println("get date failed")
	} else {
		fmt.Println("get dat succeed")
	}
}

//用Go实现向服务端页面简单的写入数据
func display_1(w http.ResponseWriter, r *http.Request) {
	fn := r.PostFormValue("first name")
	ln := r.PostFormValue("last name")

	fmt.Println("first name:" + fn)
	fmt.Println("last name:" + ln)

	w.Write([]byte("hello world"))
}

//用Go实现向get方法,在页面中动态的获取数据并且显示
/*
 * 出现问题：理解成为两个页面进行互通,出现了第二个页面获取数据出错
 */

func display_2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	arg1 := r.FormValue("arg1")
	str := fmt.Sprintf("arg1:%s", arg1)
	w.Write([]byte(str))
}
