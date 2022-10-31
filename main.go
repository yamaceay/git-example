package main

import (
	"fmt"

	"github.com/git-example/types"
)

func main() {
	comments := []types.Comment{
		{Liked: []string{"Adrian"}, Username: "Betty", Content: "nice photo!"},
	}
	posts := []types.Post{
		{Liked: []string{}, Viewed: []string{}, Tags: []string{"sunset", "beach"}},
		{Liked: []string{}, Viewed: []string{}, Location: "istanbul", Comments: comments},
	}
	stories := []types.Story{
		{Liked: []string{}, Viewed: []string{}},
		{Liked: []string{}, Viewed: []string{}},
	}
	user := types.User{
		Username:  "Adrian",
		Followers: []string{"Betty"},
		Posts:     posts,
		Stories:   stories,
	}
	fmt.Printf("%s\n", user)
}
