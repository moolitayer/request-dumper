package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		printHeaders(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9099", nil))
}

func printHeaders(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	dump := string(requestDump)
	fmt.Println(dump)
	w.Write(requestDump)
}
