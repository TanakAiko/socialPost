package main

import (
	"log"
	"net/http"
	conf "post/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//http.HandleFunc("/createPost", hd.MainHandler)

	log.Printf("Server (portAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
