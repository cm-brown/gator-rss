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
	if err = cfg.SetUser("Cameron"); err != nil {
		log.Fatal("error setting user:", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal("error reading config file: ", err)
	}
	fmt.Println(cfg)
}
