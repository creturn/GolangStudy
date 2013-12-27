package main

import (
	"fmt"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Cookie:%v", r.Cookies())
}
func StartServer() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

func main() {
	StartServer()
}
