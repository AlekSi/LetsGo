package main

import (
	"fmt"
	"net/http"
)

func main() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "今日は, Ульяновск!")
	}
	http.ListenAndServe(":6789", http.HandlerFunc(hello))
}
