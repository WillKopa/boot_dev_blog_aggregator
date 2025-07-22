package main

import (
	"log"
	"os"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	gator_config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	// Create Florida
	gator_state := state{
		cfg: gator_config,
	}

	commands := commands{
		Command_map: map[string]func(*state, command) error{},
	}
	commands.register("login", handler_login)

	if len(os.Args) < 2 {
		log.Fatal("Too few args")
	}
	cmd_args := os.Args[2:]

	cmd := command{
		Name: os.Args[1],
		Args: cmd_args,
	}

	err = commands.run(&gator_state, cmd)

	if err != nil {
		log.Fatal("error running command ", os.Args[1], ": ", err)
	}
}
