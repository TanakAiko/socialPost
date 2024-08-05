package main

import (
	// hd "post/internals/handlers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	conf "post/config"
)

func main() {
	// http.HandleFunc("/createComment", hd.createComment)

	log.Printf("Server (portAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
