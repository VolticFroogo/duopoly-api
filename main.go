package main

import (
	"log"

	"github.com/VolticFroogo/duopoly-api/db"
	"github.com/VolticFroogo/duopoly-api/helper"
	"github.com/VolticFroogo/duopoly-api/lobby"
	"github.com/VolticFroogo/duopoly-api/ws"
	"github.com/gin-gonic/gin"
)

func main() {
	helper.Seed()

	err := db.Init()
	if err != nil {
		log.Fatalf("Error initialising DB: %s", err)
	}

	r := gin.Default()

	lobby.Init(r)

	ws.Init(r)

	err = r.Run()
	if err != nil {
		log.Fatalf("Error running Gin router: %s", err)
	}
}
