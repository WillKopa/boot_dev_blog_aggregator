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
	// Get Florida
	gator_state := get_state()
	cmds := get_commands()
	cmd := parse_command(os.Args)

	err := cmds.run(gator_state, cmd)

	if err != nil {
		log.Fatal("error running command ", os.Args[1], ": ", err)
	}
}

func get_commands() *commands {
	cmds := commands{
		Command_map: map[string]func(*state, command) error{},
	}
	cmds.register("login", handler_login)
	cmds.register("register", handler_register)
	cmds.register("reset", handler_reset)
	cmds.register("users", handler_get_users)

	return &cmds
}

func parse_command(args []string) command {
	if len(args) < 2 {
		log.Fatal("Too few args")
	}
	cmd_args := os.Args[2:]

	return command{
		Name: os.Args[1],
		Args: cmd_args,
	}
}

func get_state() *state {
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

	return &gator_state
}