package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"post/internals/tools"
	md "post/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post md.Post
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Println("received data (post): ", post)

	if err := post.CreatePost(); err != nil {
		http.Error(w, "Error while creating post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, "New post created", http.StatusCreated)
}
