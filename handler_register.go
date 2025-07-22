package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handler_register(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no name given")
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("to many arguments when registering, expected only name")
	}
	name := cmd.Args[0]

	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}


	_, err := s.db.GetUser(context.Background(), name)

	if err == nil {
		log.Fatal("User with name: ", name, " already exists")
	}

	_, err = s.db.CreateUser(context.Background(), params)

	if err != nil {
		return err
	}
	s.cfg.SetUser(name)
	fmt.Println("User has been created")
	return nil
}
