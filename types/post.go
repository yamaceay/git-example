package types

import "time"

type Post struct {
	ImgHref  string
	Likes    int
	Views    int
	Tags     []string
	Location string
	Time     time.Time
	Comments []Comment
}

type Comment struct {
	Content   string
	Username  string
	CreatedAt time.Time
	Likes     int
}

type Story struct {
	Likes int
	Views int
}

type User struct {
	Username  string
	Posts     []Post
	Followers []string
	Stories   []Story
}
