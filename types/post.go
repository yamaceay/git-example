package types

import "time"

type Post struct {
	ImgHref  string
	Likes    int
	Views    int
	Tags     []string
	Location string
	Comments []string
	Time     time.Time
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
