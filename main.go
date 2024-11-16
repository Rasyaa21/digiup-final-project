package main

import (
	"base-gin/config"
	_ "base-gin/docs"
	"base-gin/repository"
	"base-gin/rest"
	"base-gin/server"
	"base-gin/service"
	"base-gin/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Sistem Peminjaman Buku API
//	@version		1.0
//	@description	API Sistem Peminjaman Buku Tugas Akhir Telkom Digiup

//	@contact.name	Rasya
//	@contact.email	naufalrasya21907@gmail.com

//	@license.name	MIT

//	@host		localhost:3000
//	@BasePath	/v1

//	@securityDefinitions.apiKey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Bearer auth containing JWT

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg := config.NewConfig()
	storage.InitDB(cfg)
	repository.SetupRepositories()
	service.SetupServices(&cfg)

	app := server.Init(&cfg, repository.GetAccountRepo())
	rest.SetupRestHandlers(app)

	// Swagger
	if cfg.App.Mode == "debug" {
		app.GET("/swagger/*any", gin.BasicAuth(gin.Accounts{
			"foo": "bar",
		}), ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	server.Serve(app)
}
