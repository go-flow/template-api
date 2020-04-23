package main

import (
	"net/http"

	"github.com/go-flow/template-api/business"
	"github.com/go-flow/template-api/db"
	"github.com/go-flow/template-api/repositories"
	"github.com/go-flow/template-api/services"

	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/config"
	"github.com/go-flow/template-api/controllers"

	_ "github.com/go-flow/template-api/docs"
)

var (
	// Version is automatically updated on build time
	Version = "v0.0.0"
)

// @title Template API API
// @version 0.1.0
// @description REST API for Template API API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi
func main() {

	//load application configuration
	opts := config.LoadWithVersion(Version)

	//get application instance
	app := flow.NewWithOptions(opts)

	// initialize database
	db.Init(app)

	// initialize repositories
	repositories.Init(app)

	// initialize services
	services.Init(app)

	//initialize business logic
	business.Init(app)

	// initialize controllers
	controllers.Init(app)

	app.Logger.Infof("Starting Template API Version: %s", Version)
	if err := app.Serve(); err != nil && err != http.ErrServerClosed {
		app.Logger.Error(err)
	}

}
