package main

import (
	"log"
	"prompt-share-backend/api"
	"prompt-share-backend/config"
	"prompt-share-backend/database"
)

// @title Prompt Share API
// @version 1.0
// @description Prompt Share 后端 API 文档
// @host localhost:8080
// @BasePath /api
func main() {
	// load config
	config.Load()

	// init db
	database.Init()

	// init router and services
	r := api.InitRouter()

	// run
	addr := config.Cfg.Server.Addr
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("Server running at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
