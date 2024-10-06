package database

import (
	"auction-systems/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgresClient(databaseConfig *config.PostgresDatabaseConfig, logger *zap.Logger) (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Database),
	)
	if err != nil {
		return nil, err
	}

	go databasePingChecker(db, databaseConfig, logger)

	return db, nil
}

func databasePingChecker(db *sql.DB, databaseConfig *config.PostgresDatabaseConfig, logger *zap.Logger) {
	for {
		time.Sleep(time.Minute)
		err := db.Ping()
		if err != nil {
			logger.Error(fmt.Sprintf("connect pgsql db %s:%s.", databaseConfig.Host, databaseConfig.Port), zap.Error(err))
		}
	}
}
