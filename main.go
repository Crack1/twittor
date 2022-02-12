package main

import (
	"log"

	"github.com/Crack1/twittor/db"
	"github.com/Crack1/twittor/handlers"
)

func main() {
	if db.ConectionCheck() == 0 {
		log.Fatal("No DB conextion")
		return
	}
	handlers.Drivers()
}
