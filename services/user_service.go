package services

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/go-flow/template-api/services/user"
)

// UserService defines set of available operations aroud Users model
type UserService interface {
	// UserService ensures interface implementation
	UserService() string

	// GetByID returns user model based on a provided id
	GetByID(id uint64) (*models.User, error)

	// GetByEmail returns user record based on a provided email
	GetByEmail(email string) (*models.User, error)

	// GetAll returns all Users model for given paging users
	GetAll(paginator *paging.Paginator) ([]*models.User, error)

	// Create new User object in database
	Create(user *models.User) error

	// Update existing User object in database
	Update(user *models.User) error

	// Save creates or updates user model based on ID user
	Save(user *models.User) error

	// Delete User object from database
	Delete(user *models.User) error
}

// NewUserService creates new UserService implementation
func NewUserService(app *flow.App) UserService {
	return user.NewService(app)
}
