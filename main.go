package main

import (
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var server = echo.New()

func main() {
	InitDatabase()
	InitMigrations()
	InitRoutes()
	server.Logger.Fatal(server.Start(":8081"))
}
