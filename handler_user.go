package main

import (
	"errors"
	"github.com/cm-brown/gator-rss/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(command.args) = 0 {
		return errors.New("Must provide an argument with command")
	}


}
