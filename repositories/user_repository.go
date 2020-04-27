package repositories

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/go-flow/template-api/repositories/user"
)

// UserRepository defines set of available operations around Users record
type UserRepository interface {
	// UserRepository ensures interface implementation
	UserRepository() string

	// GetByID returns user record based on a provided id
	GetByID(id uint64) (*models.User, error)

	// GetByEmail returns user record based on a provided email
	GetByEmail(email string) (*models.User, error)

	// GetAll returns all Users records for given paging users
	GetAll(paginator *paging.Paginator) ([]*models.User, error)

	// Create new User record in database
	Create(user *models.User) error

	// Update existing User record in database
	Update(user *models.User) error

	// Save creates or updates user record based on ID user
	Save(user *models.User) error

	// Delete User record from database
	Delete(user *models.User) error

	// DeleteByID user record from database
	DeleteByID(id uint64) error
}

// NewUserRepository registers UserRepository interface implementation
func NewUserRepository(app *flow.App) UserRepository {
	return user.NewRepository(app)
}
