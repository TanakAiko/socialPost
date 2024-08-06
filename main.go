package main

import (
	"log"
	"net/http"
	conf "post/config"
	dbmanager "post/internals/dbManager"
	hd "post/internals/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := dbmanager.InitDB()
	if err != nil {
		log.Println("db not opening !", err)
		return
	}
	defer db.Close()
	conf.DB = db

	http.HandleFunc("/post/createPost", hd.CreatePost)
	http.HandleFunc("/post/getAllPost", hd.GetAllPost)
	http.HandleFunc("/post/getAllPost", hd.GetGroupPost)

	http.HandleFunc("/post/createComment", hd.CreateComment)
	http.HandleFunc("/post/getAllPostComment", hd.GetAllPostComment)
	http.HandleFunc("/post/getLastComment", hd.GetLastComment)
	http.HandleFunc("/post/deleteComment", hd.DeleteComment)
	http.HandleFunc("/post/commentReaction", hd.CreateCommentReaction)

	log.Printf("Server (port service) started at http://localhost:%v\n", conf.Port)
	http.ListenAndServe(":"+conf.Port, nil)
}
