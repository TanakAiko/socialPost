package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"post/internals/tools"
	md "post/models"
)

func CreateCommentReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var reaction md.Comment_reaction
	json.NewDecoder(r.Body).Decode(&reaction)

	fmt.Println("received data (comment): ", reaction)

	if err := reaction.CreateCommentReaction(); err != nil {
		http.Error(w, "Error while creating reaction : "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, "New commentReaction created", http.StatusCreated)
}
