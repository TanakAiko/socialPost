package handlers

import (
	"net/http"
	"os"
	"post/config"
	"post/internals/tools"
	md "post/models"
	"strconv"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	var comment md.Comment

	content, err := os.ReadFile("./databases/sqlRequests/deleteComment.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := config.DB.Exec(string(content), comment.Id)
	if err != nil {
		http.Error(w, "Error while deleting comment : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No comment found with ID: "+strconv.Itoa(comment.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, "Comment well deleted", http.StatusOK)
}
