package db

import (
	"fmt"
	_ "github.com/lib/pq"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Config represents the database configuration.
type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	DriverName   string
}

var (
	db *gorm.DB
)

//ConnectToDB connects to the database and returns the connection
func ConnectToDB(cfg *Config) (*gorm.DB, error) {
	var err error
	// simply avoids duplicating unnecessary connections
	if db != nil {
		return db, nil
	}

	db, err = gorm.Open(pg.Open(
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable",
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.Password,
			cfg.DatabaseName,
		)))

	return db, err
}
