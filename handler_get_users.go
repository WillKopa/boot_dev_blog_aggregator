package main

import (
	"context"
	"fmt"
	"log"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
)

func handler_get_users(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal("error reading users from DB ", err)
	}
	print_users(s, users)
	return nil
}

func print_users(s *state, users []database.User) {
	for _, user := range(users) {
		fmt.Print(" * ", user.Name)
		if s.cfg.Current_user_name == user.Name {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
}
