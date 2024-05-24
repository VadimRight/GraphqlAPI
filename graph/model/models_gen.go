// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID       string `json:"id"`
	Comment  string `json:"comment"`
	AuthorID string `json:"authorId"`
	ItemID   string `json:"itemId"`
	Author   *User  `json:"author"`
}

type CommentResponse struct {
	ID              string  `json:"id"`
	Comment         string  `json:"comment"`
	AuthorID        string  `json:"authorId"`
	PostID          string  `json:"postId"`
	ParentCommentID *string `json:"parentCommentID,omitempty"`
	Author          *User   `json:"author"`
}

type Mutation struct {
}

type Post struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	AuthorID string `json:"authorId"`
	Author   *User  `json:"author"`
}

type Query struct {
}

type Token struct {
	Token string `json:"token"`
}

type User struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Posts    []*Post    `json:"posts"`
	Comments []*Comment `json:"comments"`
}
