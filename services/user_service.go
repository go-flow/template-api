package services

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/go-flow/template-api/repositories"
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
	return &userService{}
}

// userService struct that implements userService interface
type userService struct {
	// UserRepository implementation injected by dependency injection
	UserRepository repositories.UserRepository
}

// UserService ensures interface implementation
func (svc *userService) UserService() string {
	return "userService"
}

// GetByID returns user model based on a provided id
func (svc *userService) GetByID(id uint64) (*models.User, error) {
	return svc.UserRepository.GetByID(id)
}

// GetByEmail returns user record based on a provided email
func (svc *userService) GetByEmail(email string) (*models.User, error) {
	return svc.UserRepository.GetByEmail(email)
}

// GetAll returns all Users model for given paging users
func (svc *userService) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	return svc.UserRepository.GetAll(paginator)
}

// Create new User object in database
func (svc *userService) Create(user *models.User) error {
	return svc.UserRepository.Create(user)
}

// Update existing User object in database
func (svc *userService) Update(user *models.User) error {
	return svc.UserRepository.Update(user)
}

// Save creates or updates user model based on ID user
func (svc *userService) Save(user *models.User) error {
	if user.ID > 0 {
		return svc.Update(user)
	}

	return svc.Create(user)
}

// Delete User object from database
func (svc *userService) Delete(user *models.User) error {
	return svc.UserRepository.Delete(user)
}
