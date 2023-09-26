package domain

type PostComments struct {
	PostID int    `json:"post_id"`
	Word   string `json:"word"`
	Count  int    `json:"count"`
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostRepository interface {
	Comments(postId int) ([]Comment, error)
	Posts() ([]Post, error)
}

type PostUseCase interface {
	Comments(postId int) ([]Comment, error)
	Posts() ([]Post, error)
}
