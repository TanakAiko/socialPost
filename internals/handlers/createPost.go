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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post md.Post
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Println("received data (post): ", post)

	if err := post.CreatePost(); err != nil {
		http.Error(w, "Error while creating post : "+err.Error(), http.StatusInternalServerError)
		return
	}

	if post.Privacy == "almost_private" {
		for _, userId := range post.AuthList {
			if err := createPostPermission(post.Id, userId); err != nil {
				http.Error(w, "Error while insering the authorized user : "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	tools.WriteResponse(w, "New post created", http.StatusCreated)
}

func createPostPermission(postId, permisedOneId int) error {
	tx, err := config.DB.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewPostPermission.sql")
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(string(content))
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		postId,
		permisedOneId,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}
	return err
}
