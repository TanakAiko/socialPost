package handlers

import (
	"net/http"
	"os"
	"post/config"
	"post/internals/tools"
	md "post/models"
)

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./databases/sqlRequests/getAllPost.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := config.DB.Query(string(content), "all")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []md.Post{}
	for rows.Next() {
		var post md.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.Image, &post.Content, &post.Type, &post.Privacy, &post.CreatedAt); err != nil {
			http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	tools.WriteResponse(w, posts, http.StatusOK)
}
