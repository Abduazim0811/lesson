package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id    uint
	Name  string
	Email string
	Posts []Post
}

type Post struct {
	Id        uint
	UserId    uint
	Title     string
	Content   string
	CreatedAt time.Time
}

type Comment struct{
	Id uint
	Post_id uint
	Comment_context string
}

func (u *User) CreatePost(cpdto CreatePostDTO) (*Post, error) {
	query := `INSERT INTO Posts(user_id, title, content, created_at)
				VALUES ($1, $2, $3, $4) RETURNING id, user_id, title, content, created_at;`
	var newPost Post
	row := cpdto.db.QueryRow(query, u.Id, cpdto.post.Title, cpdto.post.Content, time.Now())

	if err := row.Scan(&newPost.Id, &newPost.Title, &newPost.Content, &newPost.CreatedAt); err != nil {
		return nil, err
	}

	u.Posts = append(u.Posts, newPost)
	return &newPost, nil
}

///////////////////////////////////////////////////////////////////////////////////////

type CreatePostDTO struct {
	post Post
	db   *sql.DB
}
