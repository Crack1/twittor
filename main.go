package main

import (
	"github.com/Crack1/twittor/db"
	"github.com/Crack1/twittor/handlers"
	"log"
)

func main() {
	if db.ConectionCheck() == 0 {
		log.Fatal("No DB connection")
		return
	}
	handlers.Drivers()
}
