package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"post/config"
	"post/internals/tools"
	md "post/models"
)

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var postData md.Post
	json.NewDecoder(r.Body).Decode(&postData)

	content, err := os.ReadFile("./databases/sqlRequests/getAllPost.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := config.DB.Query(string(content), postData.Id, postData.Id)
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []md.Post{}
	for rows.Next() {
		var post md.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Image, &post.Content, &post.Type, &post.Privacy, &post.CreatedAt); err != nil {
			http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	tools.WriteResponse(w, posts, http.StatusOK)
}

func GetGroupPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var postData md.Post
	json.NewDecoder(r.Body).Decode(&postData)

	content, err := os.ReadFile("./databases/sqlRequests/getGroupPost.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := config.DB.Query(string(content), postData.GroupId)
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []md.Post{}
	for rows.Next() {
		var post md.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Image, &post.Content, &post.Type, &post.Privacy, &post.CreatedAt); err != nil {
			http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	tools.WriteResponse(w, posts, http.StatusOK)
}
