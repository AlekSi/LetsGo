package main

import (
	"fmt"
	"net/http"
)

func main() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "今日は, РИТ++!")
	}
	http.ListenAndServe("127.0.0.1:6789", http.HandlerFunc(hello))
}
