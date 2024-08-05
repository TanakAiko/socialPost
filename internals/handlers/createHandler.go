package handlers

import (
	"database/sql"
	"net/http"

	"post/internals/tools"
	md "post/models"
)

func createPost(w http.ResponseWriter, post md.Post, db *sql.DB) {
	if err := post.CreatePost(db); err != nil {
		http.Error(w, "Error while creating post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteResponse(w, "New post created", http.StatusCreated)
}
