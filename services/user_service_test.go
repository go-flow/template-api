package services

import (
	"errors"
	"testing"
	"time"

	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/mocks"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	service            UserService
	userRepositoryMock *mocks.UserRepositoryMock
}

//TestServiceSuite is ure repository test suite runner
func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

// SetupSuite configures suite for unit testing
func (s *ServiceSuite) SetupSuite() {
	// anything to do speciffic to setup this suite?
}

// BeforeTest is called before every test in Service Suite
// it ensures that every test is setup correctly
func (s *ServiceSuite) BeforeTest(_, _ string) {
	s.userRepositoryMock = mocks.NewUserRepositoryMock()

	s.service = &userService{
		UserRepository: s.userRepositoryMock,
	}
}

// AfterTest is executed after every test in ServiceSuite
// it ensures that all Test Suite expectations were met
func (s *ServiceSuite) AfterTest(_, _ string) {
	//s.userRepositoryMock.AssertExpectations(s.T())
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
	// test model
	model := &models.User{
		ID:        1,
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// set method call expectations for existing user
	s.userRepositoryMock.On("GetByEmail", model.Email).Return(model, nil)

	// make actual call
	user, err := s.service.GetByEmail(model.Email)

	// assert results
	assert.Nil(s.T(), err, "expected error from first call to be <nil>, got %v", err)
	assert.Equal(s.T(), model, user, "expected test model to be Equal to returned user model")
}

func (s *ServiceSuite) Test_GetAll() {
	// test model
	model := []*models.User{
		{
			ID:        1,
			FirstName: "Sedin",
			LastName:  "Dugum",
			Email:     "sedin@mop.ba",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			FirstName: "Sedin2",
			LastName:  "Dugum2",
			Email:     "sedin+2@mop.ba",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	paginator := paging.NewWithDefaults()

	// set method call expectations for existing user
	s.userRepositoryMock.On("GetAll", paginator).Return(model, nil)

	// make actual call
	user, err := s.service.GetAll(paginator)

	// assert results
	assert.Nil(s.T(), err, "expected error from first call to be <nil>, got %v", err)
	assert.Equal(s.T(), model, user, "expected test model to be Equal to returned user model")
}

func (s *ServiceSuite) Test_Create() {
	user := &models.User{
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
	}
	// test model
	model := &models.User{
		ID:        1,
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
	}

	// set method call expectations for existing user
	s.userRepositoryMock.On("Create", user).Return(nil).Run(func(args mock.Arguments) {
		// set model ID to be the same as model object
		u := args.Get(0).(*models.User)
		u.ID = model.ID
	})

	// make actual call
	err := s.service.Save(user)

	// assert results
	assert.Nil(s.T(), err, "expected error from first call to be <nil>, got %v", err)
	assert.Equal(s.T(), model, user, "expected test model to be Equal to returned user model")
}

func (s *ServiceSuite) Test_Update() {
	user := &models.User{
		ID:        1,
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
	}
	// test model
	model := &models.User{
		ID:        1,
		FirstName: "Sedin1",
		LastName:  "Dugum1",
		Email:     "sedin@mop.ba",
	}

	// set method call expectations for existing user
	s.userRepositoryMock.On("Update", user).Return(nil).Run(func(args mock.Arguments) {
		// set model ID to be the same as model object
		u := args.Get(0).(*models.User)
		u.FirstName = model.FirstName
		u.LastName = model.LastName
	})

	// make actual call
	err := s.service.Save(user)

	// assert results
	assert.Nil(s.T(), err, "expected error from first call to be <nil>, got %v", err)
	assert.Equal(s.T(), model, user, "expected test model to be Equal to returned user model")
}

func (s *ServiceSuite) Test_Delete() {
	// test model
	model := &models.User{
		ID:        1,
		FirstName: "Sedin",
		LastName:  "Dugum",
		Email:     "sedin@mop.ba",
	}

	// set method call expectations for existing user
	s.userRepositoryMock.On("DeleteByID", uint64(1)).Return(nil)
	//s.userRepositoryMock.On("DeleteByID", uint64(1)).Return(nil)

	// make actual call
	err := s.service.Delete(model)

	// assert results
	assert.Nil(s.T(), err, "expected error from first call to be <nil>, got %v", err)
}
