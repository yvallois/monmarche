package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Connect(user string, password string, host string, port string, database string, poolSize int, logLevel logger.LogLevel) {
	dsn := paramsToDSN(user, password, host, port, database)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logLevel),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database with user %s, and name %s. Error: %s",
			user, database, err.Error())
	}
	setConnectionPool(DB, poolSize)
}

func paramsToDSN(user string, password string, host string, port string, database string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
}

func setConnectionPool(db *gorm.DB, poolSize int) {
	rawSqlDb, err := db.DB()
	if err != nil {
		panic("Unable to set postgres connection pool")
	}
	rawSqlDb.SetMaxIdleConns(poolSize / 2)
	rawSqlDb.SetMaxOpenConns(poolSize)
}
