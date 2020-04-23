package services

import "github.com/go-flow/flow"

// Init initializes services layer
func Init(app *flow.App) {
	app.Register(NewUserService(app))
}
