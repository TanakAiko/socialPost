package main

import (
	"log"
	"net/http"
	conf "post/config"
	dbManager "post/internals/dbManager"

	hd "post/internals/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := dbManager.InitDB()
	if err != nil {
		log.Println("db not opening !", err)
		return
	}
	defer db.Close()

	conf.DB = db

	http.HandleFunc("/post/create", hd.CreatePost)

	log.Printf("Server (portAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
