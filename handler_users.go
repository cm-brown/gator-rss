package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	ctx := context.Background()
	names, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("couldn't fetch users: %w", err)
	}

	for _, value := range names {
		if value == s.cfg.Username {
			fmt.Printf("* %s (current)\n", value)
		} else {
			fmt.Printf("* %s\n", value)
		}
	}
	return nil
}
