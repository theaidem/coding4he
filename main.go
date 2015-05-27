package main

import (
	"flag"
	"log"
)

var db *Store

func init() {
	db = &Store{}
}

func main() {

	tcp_port := flag.String("tcp_port", "5555", "port to run the TCP server on")
	http_port := flag.String("http_port", "8080", "port to run the HTTP server on")
	flag.Parse()

	log.Printf("start tcp server on :%s port", *tcp_port)
	if err := startTCPServer(tcp_port); err != nil {
		log.Fatal(err)
	}

	log.Printf("start http server on :%s port", *http_port)
	if err := startHTTPServer(http_port); err != nil {
		log.Fatal(err)
	}
}
