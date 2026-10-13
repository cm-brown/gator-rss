package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cm-brown/gator-rss/internal/config"
	"github.com/cm-brown/gator-rss/internal/database"
	"github.com/google/uuid"
	"log"
	"time"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide an argument with command")
	}

	ctx := context.Background()

	// TODO: finish GetUser error check
	err := s.db.GetUser(ctx)
	if err != nil {
		log.Fatal("user does not exist")
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("Username has been set")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide an argument with command")
	} else if len(cmd.args) > 1 {
		return errors.New("must provide only one user per command")
	}

	ctx := context.Background()

	userParam := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	result, err := s.db.CreateUser(ctx, userParam)
	if err != nil {
		log.Fatal(err)

	}

	s.cfg.Username = result.Name
	fmt.Printf("user, %s, was created", result.Name)

	return nil
}
