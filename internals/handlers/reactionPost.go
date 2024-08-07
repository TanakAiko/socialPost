package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"post/config"
	"post/internals/tools"
	md "post/models"
)

func SetReactionPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reacPost md.Post_reaction
	json.NewDecoder(r.Body).Decode(&reacPost)

	fmt.Println("received data (reactionPost): ", reacPost)

	tx, err := config.DB.Begin()
	if err != nil {
		log.Println(err)
		http.Error(w, "Reaction post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewPostReaction.sql")
	if err != nil {
		http.Error(w, "Reaction post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare(string(content))
	if err != nil {
		log.Println(err)
		http.Error(w, "Reaction post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		reacPost.PostId,
		reacPost.UserId,
		reacPost.Reaction,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Reaction post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		http.Error(w, "Reaction post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteResponse(w, "New reaction post created", http.StatusCreated)
}

func GetAllPostReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reacPostData md.Post_reaction
	json.NewDecoder(r.Body).Decode(&reacPostData)

	content, err := os.ReadFile("./databases/sqlRequests/getAllPostReaction.sql")
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := config.DB.Query(string(content), reacPostData.UserId, reacPostData.UserId)
	if err != nil {
		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	reacPosts := []md.Post_reaction{}
	for rows.Next() {
		var reacPost md.Post_reaction
		if err := rows.Scan(&reacPost.Id, &reacPost.PostId, &reacPost.UserId, &reacPost.Reaction); err != nil {
			http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		reacPosts = append(reacPosts, reacPost)
	}
	tools.WriteResponse(w, reacPosts, http.StatusOK)
}

func GetGroupPostReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var postData md.Post
	json.NewDecoder(r.Body).Decode(&postData)

	content, err := os.ReadFile("./databases/sqlRequests/getGroupPostReaction.sql")
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

	reacPosts := []md.Post_reaction{}
	for rows.Next() {
		var reacPost md.Post_reaction
		if err := rows.Scan(&reacPost.Id, &reacPost.PostId, &reacPost.UserId, &reacPost.Reaction); err != nil {
			http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
			return
		}
		reacPosts = append(reacPosts, reacPost)
	}
	tools.WriteResponse(w, reacPosts, http.StatusOK)
}
