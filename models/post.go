package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"
)

type Post struct {
	Id         int       `json:"postID"`
	UserId     int       `json:"userID"`
	Nickname   string    `json:"nickname"`
	Categorie  []string  `json:"categorie"`
	LikedBy    []string  `json:"likedBy"`
	DisLikedBy []string  `json:"dislikedBy"`
	Content    string    `json:"content"`
	Img        string    `json:"img"`
	NbrLike    int       `json:"nbrLike"`
	NbrDislike int       `json:"nbrDislike"`
	CreateAt   time.Time `json:"createAt"`
}

func (post *Post) CreatePost(db *sql.DB) error {
	categorieJSON, err := json.Marshal(post.Categorie)
	if err != nil {
		return err
	}

	likedByJSON, err := json.Marshal(post.LikedBy)
	if err != nil {
		return err
	}

	dislikedByJSON, err := json.Marshal(post.DisLikedBy)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
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
		post.Nickname,
		string(categorieJSON),
		string(likedByJSON),
		string(dislikedByJSON),
		post.Content,
		post.Img,
		post.NbrLike,
		post.NbrDislike,
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

func (post *Post) GetOnePost(db *sql.DB) error {
	var categorieJSON string
	var likedByJSON string
	var dislikedByJSON string
	err := db.QueryRow("SELECT userId, nickname, categorie, likedBy, dislikedBy, content, img, nbrLike, nbrDislike, createdAt FROM posts WHERE id = ?", post.Id).Scan(
		&post.UserId,
		&post.Nickname,
		&categorieJSON,
		&likedByJSON,
		&dislikedByJSON,
		&post.Content,
		&post.Img,
		&post.NbrLike,
		&post.NbrDislike,
		&post.CreateAt,
	)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(categorieJSON), &post.Categorie); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(dislikedByJSON), &post.DisLikedBy); err != nil {
		return err
	}

	err = json.Unmarshal([]byte(likedByJSON), &post.LikedBy)

	return err
}
