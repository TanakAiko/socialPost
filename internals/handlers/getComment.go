package handlers

import (
	"fmt"
	"net/http"
	"os"
	"post/config"
	"post/internals/tools"

	md "post/models"
)

func GetLastComment(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./databases/sqlRequests/getLastComment.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	var comment md.Comment
	rows, err := config.DB.Query(string(content))
	if err != nil {
		http.Error(w, "Error while getting comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	defer func() {
		if err := rows.Close(); err != nil {
			http.Error(w, "Error while closing rows: "+err.Error(), http.StatusInternalServerError)
		}
	}()

	rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.CreatedAt)

	if err := rows.Err(); err != nil {
		http.Error(w, "Error while iterating comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, comment, http.StatusOK)
}

func GetAllPostComment(w http.ResponseWriter, r *http.Request) {

	content, err := os.ReadFile("./databases/sqlRequests/getAllPostComment.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	var commentR md.Comment
	rows, err := config.DB.Query(string(content), commentR.PostId)
	if err != nil {
		http.Error(w, "Error while getting comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	defer func() {
		if err := rows.Close(); err != nil {
			http.Error(w, "Error while closing rows: "+err.Error(), http.StatusInternalServerError)
		}
	}()

	var comments []md.Comment
	for rows.Next() {
		var comment md.Comment
		if err := rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.Image, &comment.CreatedAt); err != nil {
			fmt.Println("ERROR 1")
			http.Error(w, "Error while scanning comments: "+err.Error(), http.StatusInternalServerError)
			return
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("ERROR 4")
		http.Error(w, "Error while iterating comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, comments, http.StatusOK)
}
