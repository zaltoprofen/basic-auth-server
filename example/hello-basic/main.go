package main

import (
	"fmt"
	"net/http"

	"github.com/zaltoprofen/basic-auth-server"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/", basicauth.BasicAuth(HelloHandler, func(userName, password string) bool {
		return userName == "john" && password == "password"
	}))
	http.ListenAndServe(":8080", nil)
}
