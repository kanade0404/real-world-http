package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</bpdy></html>")
}
func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listing :18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	log.Println(err)
}