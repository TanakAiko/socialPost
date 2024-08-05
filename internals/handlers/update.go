package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"post/internals/tools"
	md "post/models"
	"strconv"
)

func updateLike(w http.ResponseWriter, post md.Post, db *sql.DB) {
	query := `
        UPDATE posts
        SET nbrLike = ?, nbrDislike = ?, likedBy = ?, dislikedBy = ?
        WHERE id = ?;
    `

	likedByJSON, err := json.Marshal(post.LikedBy)
	if err != nil {
		fmt.Println("ERROR 0")
		http.Error(w, "Error while deleting post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	dislikedByJSON, err := json.Marshal(post.DisLikedBy)
	if err != nil {
		fmt.Println("ERROR 0.5")
		http.Error(w, "Error while deleting post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := db.Exec(query, post.NbrLike, post.NbrDislike, string(likedByJSON), string(dislikedByJSON), post.Id)
	if err != nil {
		fmt.Println("ERROR 1")
		http.Error(w, "Error while deleting post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("ERROR 2")
		http.Error(w, "Error while checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("ERROR 3")
		http.Error(w, "No post found with ID: "+strconv.Itoa(post.Id), http.StatusBadRequest)
		return
	}

	tools.WriteResponse(w, "Post well updated", http.StatusOK)
}
