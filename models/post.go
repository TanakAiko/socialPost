package models

import "time"

type Post struct {
	Id        int
	UserId    int
	Image     string
	Content   string
	Privacy   string
	CreatedAt time.Time
}
