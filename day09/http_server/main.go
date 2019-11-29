package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	str := "Hello"
	params := r.URL.Query()
	fnum := params.Get("fnum")
	date := params.Get("date")
	fmt.Println(fnum, date)
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/test/", test)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
