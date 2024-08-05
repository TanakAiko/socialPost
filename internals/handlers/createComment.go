package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"post/internals/tools"
	md "post/models"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var comment md.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	fmt.Println("received data (comment): ", comment)

	if err := comment.CreateComment(); err != nil {
		http.Error(w, "Error while creating comment : "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, "New comment created", http.StatusCreated)
}
