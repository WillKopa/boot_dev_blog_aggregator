package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/WillKopa/boot_dev_blog_aggregator/internal/config"
	"github.com/WillKopa/boot_dev_blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db	*database.Queries
	cfg *config.Config
}

func main() {
	gator_config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", gator_config.DB_url)
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	dbQueries := database.New(db)

	// Create Florida
	gator_state := state{
		cfg: gator_config,
		db: dbQueries,
	}

	commands := commands{
		Command_map: map[string]func(*state, command) error{},
	}

	commands.register("login", handler_login)
	commands.register("register", handler_register)
	commands.register("reset", handler_reset)

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
