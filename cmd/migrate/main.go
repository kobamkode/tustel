package main

import (
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kobamkode/tustel/internal/env"
)

var migrateType string

func init() {
	flag.StringVar(&migrateType, "run", "up", "run the migration up|down")
}

func main() {
	flag.Parse()
	env.Load()

	m, err := migrate.New("file://db/migrations", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("err: " + err.Error())
	}

	if migrateType == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err.Error())
		}
	}

	if migrateType == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err.Error())
		}
	}
}
