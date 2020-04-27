package user

import (
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/stretchr/testify/mock"
)

// NewMockRepository creates new UserRepository mocked  implementation
func NewMockRepository() *RepositoryMock {
	return &RepositoryMock{}
}

// RepositoryMock is a mocked object that implements repositories.UserRepository interface
// that describes an object that the code I am testing relies on.
type RepositoryMock struct {
	mock.Mock
}

// UserRepository ensures interface implementation
func (repo *RepositoryMock) UserRepository() string {
	args := repo.Called()
	return args.String(0)
}

// GetByID returns user model based on a provided id
func (repo *RepositoryMock) GetByID(id uint64) (*models.User, error) {
	args := repo.Called(id)

	model := args.Get(0)
	err := args.Error(1)

	if model == nil {
		return nil, err
	}
	return model.(*models.User), err
}

// GetByEmail returns user model based on a provided email
func (repo *RepositoryMock) GetByEmail(email string) (*models.User, error) {
	args := repo.Called(email)

	model := args.Get(0)
	err := args.Error(1)

	if model == nil {
		return nil, err
	}

	return model.(*models.User), err
}

// GetAll returns all Users model for given paging users
func (repo *RepositoryMock) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	args := repo.Called(paginator)

	model := args.Get(0)
	err := args.Error(1)
	if model == nil {
		return nil, err
	}
	return model.([]*models.User), err
}

// Create new User object in database
func (repo *RepositoryMock) Create(user *models.User) error {
	args := repo.Called(user)

	return args.Error(0)
}

// Update existing User object in database
func (repo *RepositoryMock) Update(user *models.User) error {
	args := repo.Called(user)

	return args.Error(0)
}

// Save creates or updates user model based on ID user
func (repo *RepositoryMock) Save(user *models.User) error {
	if user.ID > 0 {
		return repo.Update(user)
	}

	return repo.Create(user)
}

// Delete user record from database
func (repo *RepositoryMock) Delete(user *models.User) error {
	return repo.DeleteByID(user.ID)
}

// DeleteByID user record from database
func (repo *RepositoryMock) DeleteByID(id uint64) error {
	args := repo.Called(id)

	return args.Error(0)
}
