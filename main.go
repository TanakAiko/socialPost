package main

import (
	"log"
	"net/http"
	conf "post/config"
	hd "post/internals/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", hd.MainHandler)

	log.Printf("Server (portAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
