package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/cm-brown/gator-rss/internal/config"
	"github.com/cm-brown/gator-rss/internal/database"
	"github.com/google/uuid"
	"os"
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

	_, err := s.db.GetUser(ctx, cmd.args[0])
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
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
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(result.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %s\n", err)
	}
	fmt.Printf("user, %s, was created", result.Name)

	return nil
}
