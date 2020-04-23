package controllers

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/pkg/cors"
	"github.com/go-flow/template-api/pkg/swagger"

	swaggerFiles "github.com/swaggo/files"
)

// Init initializes http handlers
func Init(app *flow.App) {

	// use CORS middleware
	app.Use(cors.Default())

	// Register Controller to Flow application
	app.RegisterController(new(IndexController))
	app.RegisterController(new(UsersController))

	if app.Env == "development" {
		//init Swagger
		app.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	}

	// expose documentation to ambassador edge API documentation
	app.GET("/.ambassador-internal/openapi-docs", swagger.WrapHandler(swaggerFiles.Handler))

}
