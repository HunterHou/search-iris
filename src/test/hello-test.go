package test

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprint(w, "be")
}

func main_hello() {
	http.HandleFunc("/", helloHandler) //设置访问的路由
	err := http.ListenAndServe(":8080", nil)
	if err == nil {
		log.Fatal("ListenAndServe", err)
	}
}
