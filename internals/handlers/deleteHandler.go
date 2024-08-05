package handlers

import (
	"database/sql"
	"net/http"
	"post/internals/tools"
	md "post/models"
	"strconv"
)

func deletePost(w http.ResponseWriter, post md.Post, db *sql.DB) {
	result, err := db.Exec("DELETE FROM posts WHERE id = ?", post.Id)
	if err != nil {
		http.Error(w, "Error while deleting post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No post found with ID: "+strconv.Itoa(post.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, "Post well deleted", http.StatusOK)
}
