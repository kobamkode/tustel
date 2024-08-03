package main

import (
	"github.com/kobamkode/tustel/internal/database"
	"github.com/kobamkode/tustel/internal/env"
	"github.com/kobamkode/tustel/internal/server"
)

func init() {
	env.Load()
}

func main() {
	pool := database.NewDBConn()
	defer pool.Close()

	server.Run(pool)
}
