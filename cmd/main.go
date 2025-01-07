package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "SlideGoose test")
		})
	http.ListenAndServe(":8888", nil)
	fmt.Println("running")
}
