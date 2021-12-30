package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"pg_slowest/app/api"
	"pg_slowest/app/domain/implementation"
	"pg_slowest/app/repository/postgresql"

	_ "github.com/lib/pq"
	"pg_slowest/db"
)

func main() {
	cfg := &db.Config{
		Host:         "172.17.0.2",
		Port:         5432,
		Username:     "postgres",
		Password:     "",
		DatabaseName: "public",
		DriverName:   "postgres",
	}

	d, err := db.ConnectToDB(cfg)
	if err != nil {
		panic(err)
	}

	router := fiber.New()
	repo := postgresql.NewPGSQL(d)
	domain := implementation.NewDomain(repo)

	handler := api.NewHandler(router, domain)
	app := api.Register(handler)

	log.Fatalln(app.Listen(":8080"))
}

// CREATE EXTENSION
// SELECT total_time, query FROM pg_stat_statements ORDER BY total_time DESC LIMIT 10;
