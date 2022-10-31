package types

import (
	"encoding/json"
	"fmt"
	"time"
)

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

func (p Post) String() string {
	var bio string
	if len(p.ImgHref) == 0 {
		return bio
	}
	return jsonify(p)
}

func (s Story) String() string {
	return jsonify(s)
}

func (u User) String() string {
	bio := fmt.Sprintf("Name: %s\n", u.Username)
	if len(u.Posts) > 0 {
		bio += fmt.Sprintf("Posts: %s\n", u.Posts)
	}
	if len(u.Stories) > 0 {
		bio += fmt.Sprintf("Stories: %s\n", u.Stories)
	}
	if nfollowers := len(u.Followers); nfollowers == 0 {
		bio += fmt.Sprintf("User has %d followers!\n", nfollowers)
	} else {
		bio += fmt.Sprintf("User has no followers!\n")
	}
	return bio
}

func jsonify(v interface{}) string {
	vBytes, _ := json.MarshalIndent(v, "", "  ")
	return string(vBytes)
}
