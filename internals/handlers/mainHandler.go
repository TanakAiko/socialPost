package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	dbManager "post/internals/dbManager"
	md "post/models"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbManager.InitDB()
	if err != nil {
		log.Println("db not opening !", err)
		http.Error(w, "database can't be opened", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var req md.Request
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("req.Action : ", req.Action)

	switch req.Action {
	case "createPost":
		createPost(w, req.Body, db)
	case "getOne":
		getOnePost(w, req.Body, db)
	case "delete":
		deletePost(w, req.Body, db)
	case "getAll":
		getAllPost(w, db)
	case "updateLike":
		updateLike(w, req.Body, db)
	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

}
