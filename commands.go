package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredcommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.registeredcommands[name]; ok {
		return errors.New("command already exists")
	}

	c.registeredcommands[name] = f

	return nil
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredcommands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}

	return handler(s, cmd)
}
