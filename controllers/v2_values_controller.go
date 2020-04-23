package controllers

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/business"
)

// V2ValuesController -
type V2ValuesController struct {
	BaseController

	// ValuesBusiness implementation injected by dependency injection
	ValuesBusiness *business.ValuesBusiness
}

// Init initialize controller
func (ctrl *V2ValuesController) Init(app *flow.App) {
	ctrl.BaseController.Init(app)
}

// Version -
func (ctrl *V2ValuesController) Version() string {
	return "v2"
}

// Routes returns controller routing definition
func (ctrl *V2ValuesController) Routes() *flow.Router {
	r := flow.NewRouter()
	r.GET("/", ctrl.IndexGetAction)
	return r
}

// IndexGetAction returns list of values
// @Summary This action returns list of values from data store
// @Produce json
// @Tags index
// @Success 200 {object} models.PaginatedModel
// @Failure 400 {object} models.ResponseError
// @Router / [get]
func (ctrl *V2ValuesController) IndexGetAction(ctx *flow.Context) {
	ctrl.RenderSuccess(ctx, flow.VM{
		"message": "Hello!",
	})
}
