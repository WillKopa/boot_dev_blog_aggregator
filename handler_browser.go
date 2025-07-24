package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func handler_browse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("error parsing int from command %v", err)
		}
	}
	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("error getting posts for user: %v", err)
	}

	for _, post := range(posts) {
		print_post(post)
	}
	return nil
}

func print_post(post database.Post) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Post ID:               %v\n", post.ID)
	fmt.Printf("Post Creation Date:    %v\n", post.CreatedAt)
	fmt.Printf("Post Last Updated At:  %v\n", post.UpdatedAt)
	fmt.Printf("Post Publication Date: %v\n", post.PublishedAt)
	fmt.Printf("Post URL:              %v\n", post.Url)
	fmt.Printf("Post Title:            %v\n", post.Title)
	fmt.Printf("Post Description:      %v\n", post.Description)
	fmt.Printf("Post Feed ID:          %v\n", post.FeedID)
}