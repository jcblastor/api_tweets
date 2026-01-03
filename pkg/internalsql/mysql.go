package internalsql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jcblastor/api_tweets/internal/config"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, "America%2fArgentina%2fBuenos_Aires",
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("database is running")
	return db, nil
}
