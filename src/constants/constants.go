package constants

type ENV []string

// Env is a list of environment variables
var Env = ENV{
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",

	"PORT",
}
