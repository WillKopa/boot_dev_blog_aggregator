package main

import (
	"context"
	"fmt"
	"time"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handler_add_feed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("no name and url were not given when adding feed")
	} else if len(cmd.Args) > 2 {
		return fmt.Errorf("to many arguments for adding feed, expected only name and url")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	params := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}
	
	handler_follow(
		s, 
		command{
			Name: cmd.Name,
			Args: []string{url},
		}, 
		user,
	)

	fmt.Println(feed)
	return nil
}
