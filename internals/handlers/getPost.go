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

	post, err := getAllPost(postData.UserId)
	if err != nil {
		return
	}

	reacPost, err := getAllPostReaction(postData.UserId)
	if err != nil {
		http.Error(w, "Error while getting all post reaction : "+err.Error(), http.StatusInternalServerError)
		return
	}

	postNreac := md.PostNReac{
		Posts: post,
		Reacs: reacPost,
	}

	tools.WriteResponse(w, postNreac, http.StatusOK)
}

func GetGroupPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var postData md.Post
	json.NewDecoder(r.Body).Decode(&postData)

	post, err := getGroupPost(postData.GroupId)
	if err != nil {
		http.Error(w, "Error while getting group post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	reacPost, err := getGroupPostReaction(postData.GroupId)
	if err != nil {
		http.Error(w, "Error while getting group post reaction : "+err.Error(), http.StatusInternalServerError)
		return
	}

	postNreac := md.PostNReac{
		Posts: post,
		Reacs: reacPost,
	}

	tools.WriteResponse(w, postNreac, http.StatusOK)
}

// fonction for the GetAllPost
func getAllPost(userId int) ([]md.Post, error) {
	posts := []md.Post{}

	content, err := os.ReadFile("./databases/sqlRequests/getAllPost.sql")
	if err != nil {
		return posts, err
	}

	rows, err := config.DB.Query(string(content), userId, userId)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post md.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Image, &post.Content, &post.Type, &post.Privacy, &post.CreatedAt); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// fonction for the GetGroupPost
func getGroupPost(groupId int) ([]md.Post, error) {
	posts := []md.Post{}

	content, err := os.ReadFile("./databases/sqlRequests/getGroupPost.sql")
	if err != nil {
		return posts, err
	}

	rows, err := config.DB.Query(string(content), groupId)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post md.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Image, &post.Content, &post.Type, &post.Privacy, &post.CreatedAt); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
