package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-flow/template-api/config"
	"github.com/go-flow/template-api/models"

	"github.com/go-flow/flow"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// BaseController is base for all application controllers
// and contains common functionalities and helpers for other controllers
type BaseController struct {
	AppConfig config.AppConfig
}

// Init initialize BaseController
func (ctrl *BaseController) Init(app *flow.App) {
	ctrl.AppConfig = app.AppConfig.(config.AppConfig)
}

// Router creates application router instance with default middlewares
func (ctrl *BaseController) Router() *flow.Router {
	r := flow.NewRouter()

	return r
}

// Render renders JSON response for given parameters
func (ctrl *BaseController) Render(success bool, status int, ctx *flow.Context, data interface{}, err interface{}) {
	ctx.Status(status)
	ctx.JSON(status, models.Response{
		Success:   success,
		RequestID: ctx.RequestID(),
		Data:      data,
		Error:     err,
	})
}

// RenderSuccess renders JSON response with http status 200
func (ctrl *BaseController) RenderSuccess(ctx *flow.Context, data interface{}) {
	ctrl.Render(true, http.StatusOK, ctx, data, nil)
}

// RenderError renders Error message to JSON response
func (ctrl *BaseController) RenderError(ctx *flow.Context, code int, err error) {
	trace := ""
	if ctx.AppOptions().Env != "production" {
		te := errors.WithStack(err)
		trace = fmt.Sprintf("%+v", te)
	}

	vm := models.ResponseError{
		Cause:   errors.Cause(err).Error(),
		Message: err.Error(),
		Stack:   trace,
	}

	// check if httpError is caused by validation
	if verrs, ok := errors.Cause(err).(validator.ValidationErrors); ok {
		m := map[string]string{}

		for _, verr := range verrs {
			m[verr.Field()] = fmt.Sprintf("%s_%s", strings.ToLower(verr.Field()), verr.Tag())
		}

		vm.Validation = m
	}

	ctrl.Render(false, code, ctx, nil, vm)
}

// GenericErrorAction -
func (ctrl *BaseController) GenericErrorAction(ctx *flow.Context) {
	err := ctx.Errors.Last()
	switch ctx.Response.Status() {
	case http.StatusForbidden:
		ctrl.RenderForbiddenError(ctx, err)
	case http.StatusBadRequest:
		ctrl.RenderBadRequestError(ctx, err)
	case http.StatusMethodNotAllowed:
		ctrl.RenderNotAllowed(ctx, err)
	case http.StatusUnauthorized:
		ctrl.RenderUnauthorizedError(ctx, err)
	case http.StatusNotFound:
		ctrl.RenderNotFoundError(ctx, err)
	default:
		ctrl.RenderInternalServerError(ctx, err)
	}
}

// RenderAccepted renders JSON response with http status 202
func (ctrl *BaseController) RenderAccepted(ctx *flow.Context, data interface{}) {
	ctrl.Render(true, http.StatusAccepted, ctx, data, nil)
}

// RenderCreated renders JSON response with http status 201
func (ctrl *BaseController) RenderCreated(ctx *flow.Context, data interface{}) {
	ctrl.Render(true, http.StatusCreated, ctx, data, nil)
}

// RenderBadRequestError renders JSON response with http status 400
func (ctrl *BaseController) RenderBadRequestError(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusBadRequest, err)
}

// RenderUnauthorizedError renders JSON response with http status 401
func (ctrl *BaseController) RenderUnauthorizedError(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusUnauthorized, err)
}

// RenderPaymentRequiredError renders JSON response with http status 402
func (ctrl *BaseController) RenderPaymentRequiredError(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusPaymentRequired, err)
}

// RenderForbiddenError renders JSON response with http status 403
func (ctrl *BaseController) RenderForbiddenError(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusForbidden, err)
}

// RenderNotFoundError renders JSON response with http status 404
func (ctrl *BaseController) RenderNotFoundError(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusNotFound, err)
}

// RenderNotAllowed renders JSON response with http status 405
func (ctrl *BaseController) RenderNotAllowed(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusMethodNotAllowed, err)
}

// RenderInternalServerError renders JSON response with http status 500
func (ctrl *BaseController) RenderInternalServerError(ctx *flow.Context, err error) {
	ctrl.RenderError(ctx, http.StatusInternalServerError, err)
}
