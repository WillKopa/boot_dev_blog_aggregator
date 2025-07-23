package main

import (
	"context"
	"fmt"
)

func handler_following(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("expected no args when requesting a list of feeds you are following")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.Current_user_name)

	if err != nil {
		return fmt.Errorf("error getting list of feeds you are following: %v", err)
	}

	print(feeds)

	return nil
}

