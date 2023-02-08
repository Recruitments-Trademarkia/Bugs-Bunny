package migrations

import (
	"Bugs-Bunny/src/db"
)

// Migrate gets the gorm db connection and migrates the models present inside Package models
// models.User and models.ApiKey need to be migrated
// models.BaseModel contains uuid.UUID which is an extension to postgreSQL
//
// HINT: Check https://gorm.io/
func Migrate() {
	// TODO: Implement Migration of the database from package models

	database := db.GetDB()
	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
}
