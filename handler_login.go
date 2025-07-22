package main

import "fmt"

func handler_login(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username given when logging in")
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("to many arguments when logging in, expected only username")
	}
	err := s.cfg.SetUser(cmd.Args[0])

	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
