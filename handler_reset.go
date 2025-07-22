package main

import (
	"context"
	"fmt"
	"log"
)

func handler_reset(s *state, cmd command) error {
	err := s.db.ResetDB(context.Background())
	if err != nil {
		log.Fatal("error reseting DB ", err)
	}
	fmt.Println("Database reset success!")
	return nil
}
