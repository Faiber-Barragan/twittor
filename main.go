package main

import (
	"log"

	"github.com/Faiber-Barragan/twittor/db"
	"github.com/Faiber-Barragan/twittor/handlers"
)

func main() {
	if db.ConnectionCheck() == 0 {
		log.Fatal("No connection to the DB")
		return
	}

	handlers.Handlers()
}
