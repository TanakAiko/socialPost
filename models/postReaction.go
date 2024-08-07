package models

type Post_reaction struct {
	Id       int    `json:"postReactionId"`
	PostId   int    `json:"postId"`
	UserId   int    `json:"userId"`
	Reaction string `json:"reaction"`
}
