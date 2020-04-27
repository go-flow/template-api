package user

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/go-flow/template-api/repositories"
)

// NewService creates new Service implementation
func NewService(app *flow.App) *Service {
	return &Service{}
}

// Service struct that implements Service interface
type Service struct {
	// UserRepository implementation injected by dependency injection
	UserRepository repositories.UserRepository
}

// UserService ensures interface implementation
func (svc *Service) UserService() string {
	return "Service"
}

// GetByID returns user model based on a provided id
func (svc *Service) GetByID(id uint64) (*models.User, error) {
	return svc.UserRepository.GetByID(id)
}

// GetByEmail returns user record based on a provided email
func (svc *Service) GetByEmail(email string) (*models.User, error) {
	return svc.UserRepository.GetByEmail(email)
}

// GetAll returns all Users model for given paging users
func (svc *Service) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	return svc.UserRepository.GetAll(paginator)
}

// Create new User object in database
func (svc *Service) Create(user *models.User) error {
	return svc.UserRepository.Create(user)
}

// Update existing User object in database
func (svc *Service) Update(user *models.User) error {
	return svc.UserRepository.Update(user)
}

// Save creates or updates user model based on ID user
func (svc *Service) Save(user *models.User) error {
	if user.ID > 0 {
		return svc.Update(user)
	}

	return svc.Create(user)
}

// Delete User object from database
func (svc *Service) Delete(user *models.User) error {
	return svc.UserRepository.Delete(user)
}
