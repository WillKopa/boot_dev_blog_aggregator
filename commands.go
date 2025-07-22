package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	Command_map map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.Command_map[cmd.Name]
	if exists {
		return f(s, cmd)
	}
	return fmt.Errorf("command %v does not exist", cmd.Name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Command_map[name] = f
}
