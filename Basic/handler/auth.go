package handler

import (
	"fmt"
	"log"
	"net/http"

	
)

// Basic認証
func checkAuth(r *http.Request) bool {
	// 認証情報取得
	clientID, clientSecret, ok := r.BasicAuth()
	if ok == false {
		return false
	}
	return clientID == "client_id" && clientSecret == "client_secret"
}

func BacicAuth(w http.ResponseWriter, r *http.Request) {
	if checkAuth(r) == false {
		w.Header().Add("WWW-Authenticate", `Basic realm="SECRET AREA"`)
		w.WriteHeader(http.StatusUnauthorized) // 401
		http.Error(w, "Unauthorized", 401)
		return
	}
	_, err := fmt.Fprintf(w, "Successful Basic Authentication   \n")
	if err != nil {
		log.Fatal(err)
	}
}

