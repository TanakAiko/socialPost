package models

import (
	"log"
	"os"
	"post/config"
	"time"
)

type Post struct {
	Id        int       `json:"postId"`
	UserId    int       `json:"userId"`
	GroupId   int       `json:"groupId"`
	Image     string    `json:"image"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
	Privacy   string    `json:"privacy"`
	AuthList  []int     `json:"authList"`
	CreatedAt time.Time `json:"createdAt"`
}

func (post *Post) CreatePost() error {
	tx, err := config.DB.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewPost.sql")
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
		post.UserId,
		post.GroupId,
		post.Image,
		post.Content,
		post.Type,
		post.Privacy,
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
