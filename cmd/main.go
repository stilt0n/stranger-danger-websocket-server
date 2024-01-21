package main

import (
	"log"
	"stranger-danger/db"
	"stranger-danger/internal/user"
	ws "stranger-danger/internal/websocket"
	"stranger-danger/router"

	// Need to include postgres driver for database to work
	// this isn't used so we need to alias it w/ _
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
		return
	}
	defer dbConn.Close()
	log.Printf("Successfully connected to the database!")

	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
