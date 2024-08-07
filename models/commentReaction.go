package models

import (
	"log"
	"os"
	"post/config"
)

type Comment_reaction struct {
	Id        int    `json:"commentReactionId"`
	CommentId int    `json:"commentId"`
	UserId    int    `json:"userId"`
	Reaction  string `json:"reaction"`
}

func (commentReaction *Comment_reaction) CreateCommentReaction() error {
	tx, err := config.DB.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewCommentReaction.sql")
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
		commentReaction.CommentId,
		commentReaction.UserId,
		commentReaction.Reaction,
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
