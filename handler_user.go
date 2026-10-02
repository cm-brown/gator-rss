package main

import (
	"errors"
	"fmt"
	"github.com/cm-brown/gator-rss/internal/config"
	"github.com/cm-brown/gator-rss/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	command map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.command[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.command[name]; ok {
		return fmt.Errorf("command already exists")
	}

	c.command[name] = f

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide an argument with command")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("Username has been set")

	return nil
}

func registerUser(user string) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide an argument with command")
	} else if len(cmd.args) > 1 {
		return errors.New("must provide only one user per command")
	}

	userParam := database.CreateUserParams{
		ID:        uuid.NEW,
		CreatedAt: time.Time,
		UpdatedAt: time.Time,
		Name:      user,
	}
}
