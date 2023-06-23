package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"testproject/config"
	"testproject/internal/db"
	"testproject/internal/repository"
	"testproject/internal/server"
	"testproject/internal/service"
)

func main() {
	config.ReadConfig("./config/config.json")

	// DB connection
	db := db.ConnectionToDB()

	//Connection DB with Repo
	repo := repository.NewRepo(db)

	//Connection repository with service
	service := service.NewService(repo)

	r := gin.Default()

	//Connection service with handl
	handl := server.NewH(service, r)

	handl.AllRoutes()

	err := r.Run(config.Configure.PortRun)
	if err != nil {
		log.Fatal("router failed to start")
	}

}
