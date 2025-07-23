package main

import (
	"context"
	"fmt"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func handler_unfollow(s *state, cmd command, user database.User) error {
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return  fmt.Errorf("error, feed is not in db: %v", err)
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error, unable to unfollow: %v", err)
	}
	fmt.Println("Successfully unfollowed")
	return nil
}