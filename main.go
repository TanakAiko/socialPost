package main

import (
	// hd "post/internals/handlers"
	"log"
	"net/http"
	conf "post/config"
	dbmanager "post/internals/dbManager"
	hd "post/internals/handlers"
)

func main() {

	db,err := dbmanager.InitDB()
	if err != nil {
		log.Println("db not opening !",err)
		return
	}
	defer db.Close()
	conf.DB = db
	http.HandleFunc("/post/createPost", hd.CreatePost)
	http.HandleFunc("/post/createComment", hd.CreateComment)

	log.Printf("Server (portAPI) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
