package mocks

import (
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/stretchr/testify/mock"
)

// NewUserRepositoryMock creates new UserRepository mocked  implementation
func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

// UserRepositoryMock is a mocked object that implements repositories.UserRepository interface
// that describes an object that the code I am testing relies on.
type UserRepositoryMock struct {
	mock.Mock
}

// UserRepository ensures interface implementation
func (repo *UserRepositoryMock) UserRepository() string {
	args := repo.Called()
	return args.String(0)
}

// GetByID returns user model based on a provided id
func (repo *UserRepositoryMock) GetByID(id uint64) (*models.User, error) {
	args := repo.Called(id)

	model := args.Get(0)
	err := args.Error(1)

	if model == nil {
		return nil, err
	}
	return model.(*models.User), err
}

// GetByEmail returns user model based on a provided email
func (repo *UserRepositoryMock) GetByEmail(email string) (*models.User, error) {
	args := repo.Called(email)

	model := args.Get(0)
	err := args.Error(1)

	if model == nil {
		return nil, err
	}

	return model.(*models.User), err
}

// GetAll returns all Users model for given paging users
func (repo *UserRepositoryMock) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	args := repo.Called(paginator)

	model := args.Get(0)
	err := args.Error(1)
	if model == nil {
		return nil, err
	}
	return model.([]*models.User), err
}

// Create new User object in database
func (repo *UserRepositoryMock) Create(user *models.User) error {
	args := repo.Called(user)

	return args.Error(0)
}

// Update existing User object in database
func (repo *UserRepositoryMock) Update(user *models.User) error {
	args := repo.Called(user)

	return args.Error(0)
}

// Save creates or updates user model based on ID user
func (repo *UserRepositoryMock) Save(user *models.User) error {
	if user.ID > 0 {
		return repo.Update(user)
	}

	return repo.Create(user)
}

// Delete user record from database
func (repo *UserRepositoryMock) Delete(user *models.User) error {
	return repo.DeleteByID(user.ID)
}

// DeleteByID user record from database
func (repo *UserRepositoryMock) DeleteByID(id uint64) error {
	args := repo.Called(id)

	return args.Error(0)
}
