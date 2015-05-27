package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func startTCPServer(port *string) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		return err
	}

	go func(ln net.Listener) {
		for {
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			go handleConnection(conn)
		}
	}(ln)

	return nil
}

func startHTTPServer(port *string) error {

	http.HandleFunc("/stats", getStats)
	return http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		conn.Write([]byte(fmt.Sprintf("Oh! Error: %s.", err.Error())))
		log.Println(err)
		return
	}

	db.Add(strings.Split(string(buf), " "))
	conn.Write([]byte(fmt.Sprintf("Ok! You just wrote %d bytes.\n", len(buf))))
}

func getStats(res http.ResponseWriter, req *http.Request) {

	n := req.FormValue("N")
	var num int64 = 5
	var err error
	if len(n) != 0 {
		num, err = strconv.ParseInt(n, 0, 0)
		if err != nil {
			log.Println(err)
		}
		if num == 0 {
			num = 5
		}
	}

	stats := make(map[string]interface{})

	stats["count"] = len(db.Words.All)

	if len(db.Words.Sorted) <= int(num) {
		stats["top_"+ToStr(len(db.Words.Sorted))+"_words"] = db.Words.Sorted[:]
	} else {
		stats["top_"+ToStr(num)+"_words"] = db.Words.Sorted[:int(num)]
	}

	if len(db.Letters.Sorted) <= int(num) {
		stats["top_"+ToStr(len(db.Letters.Sorted))+"_letters"] = db.Letters.Sorted[:]
	} else {
		stats["top_"+ToStr(num)+"_letters"] = db.Letters.Sorted[:int(num)]
	}

	js, err := json.Marshal(stats)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(js)
}
