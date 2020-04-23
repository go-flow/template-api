package business

import "github.com/go-flow/flow"

// Init initializes project business layer
func Init(app *flow.App) {
	app.Register(NewUserBusiness(app))
}
