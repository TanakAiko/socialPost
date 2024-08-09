package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"post/config"
	"post/internals/tools"

	md "post/models"
)

// func GetLastComment(w http.ResponseWriter, r *http.Request) {
// 	content, err := os.ReadFile("./databases/sqlRequests/getLastComment.sql")
// 	if err != nil {
// 		http.Error(w, "Error while getting all post : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	var comment md.Comment
// 	rows, err := config.DB.Query(string(content))
// 	if err != nil {
// 		http.Error(w, "Error while getting comments: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer rows.Close()

// 	defer func() {
// 		if err := rows.Close(); err != nil {
// 			http.Error(w, "Error while closing rows: "+err.Error(), http.StatusInternalServerError)
// 		}
// 	}()

// 	rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.CreatedAt)

// 	if err := rows.Err(); err != nil {
// 		http.Error(w, "Error while iterating comments: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	tools.WriteResponse(w, comment, http.StatusOK)
// }

func GetAllPostComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var commentR md.Comment
	json.NewDecoder(r.Body).Decode(&commentR)

	content, err := os.ReadFile("./databases/sqlRequests/getAllPostComment.sql")
	if err != nil {
		http.Error(w, "Error while getting all comments : "+err.Error(), http.StatusInternalServerError)
		return
	}
	rows, err := config.DB.Query(string(content),commentR.PostId)
	if err != nil {
		http.Error(w, "Error while getting comments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

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

	tools.WriteResponse(w, comments, http.StatusOK)
	GetAllCommentReaction()
}

func CreateCommentReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var reaction md.Comment_reaction
	json.NewDecoder(r.Body).Decode(&reaction)

	fmt.Println("received data (comment): ", reaction)

	if err := reaction.CreateCommentReaction(); err != nil {
		http.Error(w, "Error while creating reaction : "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.WriteResponse(w, "New commentReaction created", http.StatusCreated)
}

func GetAllCommentReaction() []md.Comment_reaction{
	var w http.ResponseWriter
	// var r http.Request
	content, err := os.ReadFile("./databases/sqlRequests/getAllReaction.sql")
	if err != nil {
		http.Error(w, "Error while getting all reactions : "+err.Error(), http.StatusInternalServerError)
		return nil
	}

	var commentReaction md.Comment_reaction
	rows, err := config.DB.Query(string(content), commentReaction.CommentId)
	if err != nil {
		http.Error(w, "Error while getting reactions: "+err.Error(), http.StatusInternalServerError)
		return nil
	}

	defer rows.Close()

	defer func() {
		if err := rows.Close(); err != nil {
			http.Error(w, "Error while closing rows: "+err.Error(), http.StatusInternalServerError)
		}
	}()

	var commentReactions []md.Comment_reaction
	for rows.Next() {
		var comment md.Comment_reaction
		if err := rows.Scan(&comment.Id, &comment.CommentId, &comment.UserId, &comment.Reaction); err != nil {
			fmt.Println("ERROR 1")
			http.Error(w, "Error while scanning comments: "+err.Error(), http.StatusInternalServerError)
			return nil
		}

		commentReactions = append(commentReactions, comment)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("ERROR 4")
		http.Error(w, "Error while iterating comments: "+err.Error(), http.StatusInternalServerError)
		return nil
	}
	tools.WriteResponse(w, commentReactions, http.StatusOK)
	return commentReactions
}
