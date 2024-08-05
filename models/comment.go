package models

import (
	"log"
	"os"
	"post/config"
	"time"
)

type Comment struct {
	Id       int       `json:"commentId"`
	PostId   int       `json:"postId"`
	UserId   int       `json:"userId"`
	Content  string    `json:"content"`
	Image    string    `json:"image"`
	CreateAt time.Time `json:"createAt"`
}

func (comment *Comment) CreateComment() error {
	tx, err := config.DB.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewComment.sql")
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
		comment.PostId,
		comment.UserId,
		comment.Content,
		comment.Image,
		time.Now().Format(time.RFC3339),
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
