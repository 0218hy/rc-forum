package comments

type CreateCommentRequest struct {
	PostID   int32  `json:"post_id"`
	AuthorID int32  `json:"author_id"`
	Body     string `json:"body"`
}

