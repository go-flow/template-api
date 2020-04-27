package controllers

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/business"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
)

// UsersController -
type UsersController struct {
	BaseController

	// UserBusiness implementation injected by dependency injection
	UserBusiness business.UserBusiness
}

// Init initialize controller
func (ctrl *UsersController) Init(app *flow.App) {
	ctrl.BaseController.Init(app)
}

// Routes returns controller routing definition
func (ctrl *UsersController) Routes() *flow.Router {
	r := flow.NewRouter()
	r.GET("/", ctrl.IndexGetAction)
	return r
}

// IndexGetAction returns list of users
// @Summary This action returns list of users from data store
// @Produce json
// @Tags users
// @Tags index
// @Success 200 {object} models.PaginatedModel
// @Failure 400 {object} models.ResponseError
// @Router /users/ [get]
func (ctrl *UsersController) IndexGetAction(ctx *flow.Context) {
	// get paging from query string
	paginator := paging.NewPaginatorFromParams(ctx.Request.URL.Query())

	users, err := ctrl.UserBusiness.GetAll(paginator)

	if err != nil {
		ctrl.RenderInternalServerError(ctx, err)
		return
	}

	ctrl.RenderSuccess(ctx, models.PaginatedModel{
		Results:   users,
		Paginator: paginator,
	})
}
