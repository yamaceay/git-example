package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	ImgHref  string
	Liked    []string
	Viewed   []string
	Tags     []string
	Location string
	Time     time.Time
	Comments []Comment
}

type Comment struct {
	Content   string
	Username  string
	CreatedAt time.Time
	Liked     []string
}

type Story struct {
	Liked  []string
	Viewed []string
}

type User struct {
	Username  string
	Posts     []Post
	Followers []string
	Stories   []Story
}

func (u *User) Follow(u2 User) error {
	for _, follower := range u.Followers {
		if follower == u2.Username {
			return fmt.Errorf("duplicate follow request!\n")
		}
	}
	u.Followers = append(u.Followers, u2.Username)
	return nil
}

func (u *User) Like(p Post) error {
	for _, like := range p.Liked {
		if like == u.Username {
			return fmt.Errorf("duplicate follow request!\n")
		}
	}
	p.Liked = append(p.Liked, u.Username)
	return nil
}

func (p Post) String() string {
	var bio string
	if len(p.ImgHref) == 0 {
		return bio
	}
	return jsonify(
		struct {
			ImgHref  string
			Likes    int
			Views    int
			Tags     []string
			Location string
			Comments int
		}{
			ImgHref:  p.ImgHref,
			Likes:    len(p.Liked),
			Views:    len(p.Viewed),
			Tags:     p.Tags,
			Location: p.Location,
			Comments: len(p.Comments),
		})
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
	if nfollowers := len(u.Followers); nfollowers > 0 {
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
