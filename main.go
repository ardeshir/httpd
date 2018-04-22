package main

import (
	"log"
	"flag"
	"net/http"
	
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "port number")
	flag.Parse()
    http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.FileServer(http.Dir("/var/www/html/public")))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
