package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"post/internals/tools"
	md "post/models"
)

func getOnePost(w http.ResponseWriter, post md.Post, db *sql.DB) {
	if err := post.GetOnePost(db); err != nil {
		http.Error(w, "Error while getting post : "+err.Error(), http.StatusBadRequest)
		return
	}
	tools.WriteResponse(w, post, http.StatusOK)
}

func getAllPost(w http.ResponseWriter, db *sql.DB) {
	rows, err := db.Query("SELECT id, userId, nickname, categorie, likedBy, dislikedBy, content, img, nbrLike, nbrDislike, createdAt FROM posts ORDER BY createdAt DESC")
	if err != nil {
		http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []md.Post{}
	for rows.Next() {
		var post md.Post
		var categorieJSON string
		var likedByJSON string
		var dislikedByJSON string
		if err := rows.Scan(&post.Id, &post.UserId, &post.Nickname, &categorieJSON, &likedByJSON, &dislikedByJSON, &post.Content, &post.Img, &post.NbrLike, &post.NbrDislike, &post.CreateAt); err != nil {
			fmt.Println("ERROR 1")
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal([]byte(categorieJSON), &post.Categorie); err != nil {
			fmt.Println("ERROR 2")
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal([]byte(likedByJSON), &post.LikedBy); err != nil {
			fmt.Println("ERROR 3")
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal([]byte(dislikedByJSON), &post.DisLikedBy); err != nil {
			fmt.Println("ERROR 4")
			http.Error(w, "Error while getting post : "+err.Error(), http.StatusInternalServerError)
			return
		}

		posts = append(posts, post)

	}

	tools.WriteResponse(w, posts, http.StatusOK)
}
