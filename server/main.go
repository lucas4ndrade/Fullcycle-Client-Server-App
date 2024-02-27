package main

import (
	"github.com/go-fullcicle/server/config"
	"github.com/go-fullcicle/server/database"
	"github.com/go-fullcicle/server/http"
)

// StartServer starts the HTTP server
func main() {
	config.ParseConfig()
	database.StartDatabase()
	http.StartRoutes()
}
