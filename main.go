package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/git-example/types"
)

func main() {
	posts := []types.Post{
		{Likes: 300, Views: 500, Tags: []string{"sunset", "beach"}, Time: time.Now()},
		{Likes: 500, Views: 1000, Location: "istanbul", Comments: []string{"Nice pic!"}, Time: time.Now()},
	}
	stories := []types.Story{
		{Likes: 100, Views: 1000},
		{Likes: 200, Views: 300},
	}
	user := types.User{
		Followers: []string{"Adrian"},
		Posts:     posts,
		Stories:   stories,
	}

	userBytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("%s\n", string(userBytes))
}
