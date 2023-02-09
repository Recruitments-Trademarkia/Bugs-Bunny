package db

import (
	"Bugs-Bunny/src/config"
	"Bugs-Bunny/src/utils"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	// Try Until done
	try := 0
	for try < 5 {
		db, err := Connect()
		if err != nil {
			try++
			time.Sleep(time.Second * 20)
		} else {
			return db
		}
	}
	log.Panicln("Retries for DB Connection Completed")
	return nil
}

func Connect() (*gorm.DB, error) {
	utils.ImportEnv()
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=disable password=%s", config.DBConfig.Host, config.DBConfig.Username, config.DBConfig.Name, config.DBConfig.Password)
	sqlDB, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Error),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")

	return db, nil
}
