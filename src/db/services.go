package db

import (
	"Bugs-Bunny/pkg/api_keys"
	"Bugs-Bunny/pkg/user"
)

var (
	UserService user.Service     = nil
	ApiService  api_keys.Service = nil
)

func InitServices() {
	// TODO: Comment Out
	db := GetDB()

	userRepo := user.NewPostgresRepo(db)
	apiRepo := api_keys.NewPostgresRepo(db)

	UserService = user.NewService(userRepo)
	ApiService = api_keys.NewService(apiRepo)
}
