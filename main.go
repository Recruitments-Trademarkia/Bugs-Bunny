package main

import (
	"Bugs-Bunny/src/config"
	"Bugs-Bunny/src/fiber"
	"Bugs-Bunny/src/utils"
)

func main() {
	// Import Environment
	utils.ImportEnv()

	// Load Config
	config.Init()

	// Run Server
	fiber.Start()
}
