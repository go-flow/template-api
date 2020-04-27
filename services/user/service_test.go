package user

import (
	"errors"
	"testing"
	"time"

	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/repositories/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	service            *Service
	userRepositoryMock *user.RepositoryMock
}

//TestServiceSuite is ure repository test suite runner
func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

// SetupSuite configures suite for unit testing
func (s *ServiceSuite) SetupSuite() {
	s.userRepositoryMock = user.NewMockRepository()

	s.service = &Service{
		UserRepository: s.userRepositoryMock,
	}
}

// AfterTest ensures that all Test Suite expectations were met
func (s *ServiceSuite) AfterTest(_, _ string) {
	s.userRepositoryMock.AssertExpectations(s.T())
}

func (s *ServiceSuite) Test_GetByID() {
	// test model
	model := &models.User{
		ID:        1,
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	errNotFound := errors.New("user not found")

	// set method call expectations for existing user
	s.userRepositoryMock.On("GetByID", uint64(1)).Return(model, nil)
	// set method call expectations for non existing user
	s.userRepositoryMock.On("GetByID", uint64(2)).Return(nil, errNotFound)

	// make actual call
	user, err := s.service.GetByID(1)

	// assert results
	assert.Nil(s.T(), err, "expected error from first call to be <nil>, got %v", err)
	assert.Equal(s.T(), model, user, "expected test model to be Equal to returned user model")

	// make negative call
	user, err = s.service.GetByID(2)

	if assert.NotNil(s.T(), err, "expected returned error object to have value") {
		assert.Equal(s.T(), errNotFound, err, "expected err object to have `user not found` value, got %v", err)
	}
	assert.Nil(s.T(), user, "expected returned user object to be <nil>, got , %v", user)

}

func (s *ServiceSuite) Test_GetByEmail() {
}

func (s *ServiceSuite) Test_GetAll() {
}

func (s *ServiceSuite) Test_Create() {
}

func (s *ServiceSuite) Test_Update() {
}

func (s *ServiceSuite) Test_Delete() {
}
