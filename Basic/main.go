package main

import (
	"net/http"

	"github.com/seipan/Go-Authentic/tree/main/Basic/handler"
)

func main() {
	historyCreate := http.HandlerFunc(handler.BacicAuth)
	http.Handle("/api/history/create", historyCreate)

	http.ListenAndServe(":8080", nil)
}
