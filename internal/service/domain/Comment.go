package domain

type Comment struct {
	PostId int    `json:"post_id"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
