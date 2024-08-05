package handlers

import (
	"post/internals/tools"
	md "post/models"
	"database/sql"
	"fmt"
	"net/http"
)

func createComment(w http.ResponseWriter, comment md.Comment, db *sql.DB) {
	if err := comment.CreateComment(db); err != nil {
		fmt.Println("ERROR : ", err.Error())
		http.Error(w, "Error while creating comment : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tools.WriteResponse(w, "New comment created", http.StatusCreated)
}
