package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/service/c", helloHandler)
	fmt.Println("start service: C")

	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("start service err: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request header: ", r.Header)
	fmt.Fprint(w, "C, Hello, World!")
}
