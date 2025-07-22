package main

import (
	"context"
	"fmt"
	"log"
)

func handler_login(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username given when logging in")
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("to many arguments when logging in, expected only username")
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)

	if err != nil {
		log.Fatal("User ", name, " does not exist")
	}

	err = s.cfg.SetUser(name)

	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
