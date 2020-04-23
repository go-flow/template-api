package repositories

import "github.com/go-flow/flow"

// Init initializes project repositories
func Init(app *flow.App) {
	app.Register(NewUserRepository(app))
}
