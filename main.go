package main

import _ "github.com/lib/pq"

import (
	"database/sql"
	"fmt"
	"github.com/cm-brown/gator-rss/internal/config"
	"github.com/cm-brown/gator-rss/internal/database"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config file: ", err)
	}

	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		fmt.Printf("error: %s:\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	appState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	appCommands := commands{
		command: make(map[string]func(*state, command) error),
	}

	appCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("must provide a command")
		os.Exit(1)
	}

	userCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = appCommands.run(&appState, userCommand)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
