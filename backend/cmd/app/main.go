package main

import (
	_ "avito/docs"
	"avito/internal/app"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a server to assigne 2 random person from team to PR

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token. Example: "Bearer {token}"

func main() {
	app := app.NewApp()
	app.Start()
}
