package main

import (
	"flag"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "port number")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}
