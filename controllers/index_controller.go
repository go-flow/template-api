package controllers

import (
	"github.com/go-flow/flow"
)

// IndexController -
type IndexController struct {
	BaseController
}

// Init initialize controller
func (ctrl *IndexController) Init(app *flow.App) {
	ctrl.BaseController.Init(app)
	// handle 404 and 500 errors
	app.NotFoundHandler(ctrl.GenericErrorAction)
	app.ErrorHandler(ctrl.GenericErrorAction)
	app.UnauthorizedHandler(ctrl.GenericErrorAction)
}

// Routes returns controller routing definition
func (ctrl *IndexController) Routes() *flow.Router {
	r := flow.NewRouter()
	r.GET("/", ctrl.IndexGetAction)
	return r
}

// IndexGetAction greets user a Hello message
// @Summary This action does not do anything special now.
// @Produce json
// @Tags index
// @Success 200 {object} flow.VM
// @Failure 400 {object} models.ResponseError
// @Router / [get]
func (ctrl *IndexController) IndexGetAction(ctx *flow.Context) {
	ctrl.RenderSuccess(ctx, flow.VM{
		"message": "Hello!",
	})
}

// HealthGetAction checks service health
// @Summary checks if all service dependecies are working properly
// @Produce json
// @Tags health
// @Success 200 {object} flow.VM
// @Failure 500 {object} models.ResponseError
// @Router /health/ [get]
func (ctrl *IndexController) HealthGetAction(ctx *flow.Context) {

	ctrl.RenderSuccess(ctx, flow.VM{
		"health": "OK",
	})
}
