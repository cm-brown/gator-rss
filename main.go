package main

import (
	"fmt"
	"github.com/cm-brown/gator-rss/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config file: ", err)
	}

	appState := state{
		cfg: &cfg,
	}

	appCommands := commands{
		command: make(map[string]func(*state, command) error),
	}

	appCommands.register("login", handlerLogin)

	fmt.Println(appState, appCommands)
}
