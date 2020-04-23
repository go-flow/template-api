package business

import (
	"time"

	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/go-flow/template-api/services"
)

// UserBusiness defines set of business rules related to Users model
type UserBusiness interface {
	// UserBusiness ensures interface implementation
	UserBusiness() string

	// GetByID returns user model for provided id
	GetByID(id uint64) (*models.User, error)

	// GetByEmail returns user model for provided email
	GetByEmail(email string) (*models.User, error)

	// GetAll returns all Users model for given paging users
	GetAll(paginator *paging.Paginator) ([]*models.User, error)

	// Create new User record in database
	Create(firstName *string, lastName *string, profileImage *string, birthDate *time.Time, bio *string, phoneNumber *string, country *string, state *string, area *string, city *string, address *string, postCode *string) (*models.User, error)

	// Update User record in database
	Update(userID uint64, firstName *string, lastName *string, profileImage *string, birthDate *time.Time, bio *string, phoneNumber *string, country *string, state *string, area *string, city *string, address *string, postCode *string) (*models.User, error)

	// Delete User object from database
	Delete(id uint64) error
}

// NewUserBusiness creates new users business rules implementation instance
func NewUserBusiness(app *flow.App) UserBusiness {
	return &userBusiness{}
}

// userBusiness struct that implements UserBusiness interface
type userBusiness struct {
	UserService services.UserService
}

// UserBusiness ensures interface implementation
func (bl *userBusiness) UserBusiness() string {
	return "userBusiness"
}

// GetByID returns user model for provided id
func (bl *userBusiness) GetByID(id uint64) (*models.User, error) {
	return bl.UserService.GetByID(id)
}

// GetByEmail returns user model for provided email
func (bl *userBusiness) GetByEmail(email string) (*models.User, error) {
	return bl.UserService.GetByEmail(email)
}

// GetAll returns all Users model for given paging users
func (bl *userBusiness) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	return bl.UserService.GetAll(paginator)
}

// Create new User object in database
func (bl *userBusiness) Create(firstName *string, lastName *string, profileImage *string, birthDate *time.Time, bio *string, phoneNumber *string, country *string, state *string, area *string, city *string, address *string, postCode *string) (*models.User, error) {
	user := new(models.User)

	if firstName != nil {
		user.FirstName = *firstName
	}

	if lastName != nil {
		user.LastName = *lastName
	}

	if profileImage != nil {
		user.ProfileImage = *profileImage
	}

	if bio != nil {
		user.Bio = *bio
	}
	if phoneNumber != nil {
		user.PhoneNumber = *phoneNumber
	}
	if country != nil {
		user.Country = *country
	}
	if state != nil {
		user.State = *state
	}

	if area != nil {
		user.Area = *area

	}
	if city != nil {
		user.City = *city
	}
	if address != nil {
		user.Address = *address
	}
	if postCode != nil {
		user.PostCode = *postCode
	}

	user.BirthDate = birthDate

	return user, bl.UserService.Create(user)
}

// Update existing User object in database
func (bl *userBusiness) Update(userID uint64, firstName *string, lastName *string, profileImage *string, birthDate *time.Time, bio *string, phoneNumber *string, country *string, state *string, area *string, city *string, address *string, postCode *string) (*models.User, error) {
	user, err := bl.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if firstName != nil {
		user.FirstName = *firstName
	}

	if lastName != nil {
		user.LastName = *lastName
	}

	if profileImage != nil {
		user.ProfileImage = *profileImage
	}

	if bio != nil {
		user.Bio = *bio
	}
	if phoneNumber != nil {
		user.PhoneNumber = *phoneNumber
	}
	if country != nil {
		user.Country = *country
	}
	if state != nil {
		user.State = *state
	}

	if area != nil {
		user.Area = *area

	}
	if city != nil {
		user.City = *city
	}
	if address != nil {
		user.Address = *address
	}
	if postCode != nil {
		user.PostCode = *postCode
	}

	user.BirthDate = birthDate

	return user, bl.UserService.Update(user)
}

// Delete User object from database
func (bl *userBusiness) Delete(id uint64) error {
	user, err := bl.GetByID(id)
	if err != nil {
		return err
	}
	return bl.UserService.Delete(user)
}
