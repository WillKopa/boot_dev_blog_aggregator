package main

import (
	"context"
	"fmt"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
		if err != nil {
			return fmt.Errorf("error getting user id: %v", err)
		}
		return handler(s, cmd, user)
	}
}