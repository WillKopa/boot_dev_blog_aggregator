package main

import (
	"context"
	"fmt"
	"time"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handler_follow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no url given")
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("to many arguments when following, expected only url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed follow %v", err)
	}

	fmt.Printf("Feed follow: %v\n", feed_follow)

	return nil
}

