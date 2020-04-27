package user

import (
	"errors"
	"time"

	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/stretchr/testify/mock"
)

// NewMockRepository creates new UserRepository mocked  implementation
func NewMockRepository(testData []*models.User) *RepositoryMock {
	return &RepositoryMock{
		data: testData,
	}
}

// RepositoryMock is a mocked object that implements repositories.UserRepository interface
// that describes an object that the code I am testing relies on.
type RepositoryMock struct {
	mock.Mock

	data []*models.User
}

// UserRepository ensures interface implementation
func (repo *RepositoryMock) UserRepository() string {
	repo.Called()
	return "RepositoryMock"
}

// GetByID returns user model based on a provided id
func (repo *RepositoryMock) GetByID(id uint64) (*models.User, error) {
	repo.Called(id)
	var model *models.User

	for _, val := range repo.data {
		if val.ID == id {
			model = val
		}
	}
	if model == nil {
		return nil, errors.New("record not found")
	}
	return model, nil
}

// GetByEmail returns user model based on a provided email
func (repo *RepositoryMock) GetByEmail(email string) (*models.User, error) {
	repo.Called(email)
	var model *models.User

	for _, val := range repo.data {
		if val.Email == email {
			model = val
		}
	}
	if model == nil {
		return nil, errors.New("record not found")
	}
	return model, nil
}

// GetAll returns all Users model for given paging users
func (repo *RepositoryMock) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	repo.Called(paginator)
	size := len(repo.data)
	data := []*models.User{}
	if paginator.PerPage < size {
		data = repo.data[0:paginator.PerPage]
	}

	paginator.CurrentEntriesSize = len(data)
	paginator.TotalEntriesSize = size
	paginator.TotalPages = paginator.TotalEntriesSize / paginator.PerPage
	if paginator.TotalEntriesSize%paginator.PerPage > 0 {
		paginator.TotalPages = paginator.TotalPages + 1
	}

	return data, nil
}

// Create new User object in database
func (repo *RepositoryMock) Create(user *models.User) error {
	repo.Called(user)
	user.ID = uint64(len(repo.data))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	repo.data = append(repo.data, user)
	return nil
}

// Update existing User object in database
func (repo *RepositoryMock) Update(user *models.User) error {
	repo.Called(user)
	user.UpdatedAt = time.Now()
	return nil
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
	repo.Called(id)
	var model *models.User

	for _, val := range repo.data {
		if val.ID == id {
			model = val
		}
	}
	if model == nil {
		return errors.New("record not found")
	}
	return nil
}
