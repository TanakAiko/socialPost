package models

type Request struct {
	Action string `json:"action"`
	Body   Post   `json:"body"`
}
