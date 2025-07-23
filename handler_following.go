package main

import (
	"context"
	"fmt"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func handler_following(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("expected no args when requesting a list of feeds you are following")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("error getting list of feeds you are following: %v", err)
	}

	print_feed_follows(feeds)

	return nil
}

func print_feed_follows(feeds []database.GetFeedFollowsForUserRow) {
	fmt.Printf("You are following\n")
	for _, feed_follow := range(feeds) {
		fmt.Printf(" * %s\n", feed_follow.FeedName)
	}
}

