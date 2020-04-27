package repositories

import (
	"fmt"

	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/db"
	"github.com/go-flow/template-api/models"
	"github.com/go-flow/template-api/pkg/paging"
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

// NewUserRepository creates new UserRepository implementation
func NewUserRepository(app *flow.App) UserRepository {
	return &userRepository{}
}

// userRepository struct that implements UserRepository interface
type userRepository struct {
	Store db.Store
}

// UserRepository ensures interface implementation
func (repo *userRepository) UserRepository() string {
	return "userRepository"
}

// GetByID returns user model based on a provided id
func (repo *userRepository) GetByID(id uint64) (*models.User, error) {
	model := new(models.User)

	tx := repo.Store.Where("id = ?", id).First(model)

	return model, tx.Error
}

// GetByEmail returns user model based on a provided email
func (repo *userRepository) GetByEmail(email string) (*models.User, error) {
	model := new(models.User)

	tx := repo.Store.Where("email = ?", email).First(model)

	return model, tx.Error
}

// GetAll returns all Users model for given paging users
func (repo *userRepository) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
	model := make([]*models.User, 0)

	//construct base query
	query := repo.Store.
		Limit(paginator.PerPage).
		Offset(paginator.Offset).
		Order(paginator.Order("id"))

		// apply filtering
	if len(paginator.Filter) > 2 {
		filter := fmt.Sprintf("%%%s", paginator.Filter)
		query = query.Where("first_name LIKE ?", filter).Or("last_name LIKE ?", filter)
	}

	tx := query.Find(&model).
		Offset(0).Limit(-1).Count(&paginator.TotalEntriesSize)

	paginator.CurrentEntriesSize = len(model)
	paginator.TotalPages = paginator.TotalEntriesSize / paginator.PerPage
	if paginator.TotalEntriesSize%paginator.PerPage > 0 {
		paginator.TotalPages = paginator.TotalPages + 1
	}

	return model, tx.Error
}

// Create new User object in database
func (repo *userRepository) Create(user *models.User) error {
	tx := repo.Store.Create(user)
	return tx.Error
}

// Update existing User object in database
func (repo *userRepository) Update(user *models.User) error {
	tx := repo.Store.Save(user)
	return tx.Error
}

// Save creates or updates user model based on ID user
func (repo *userRepository) Save(user *models.User) error {
	if user.ID > 0 {
		return repo.Update(user)
	}
	return repo.Create(user)
}

// Delete user record from database
func (repo *userRepository) Delete(user *models.User) error {
	return repo.DeleteByID(user.ID)
}

// DeleteByID user record from database
func (repo *userRepository) DeleteByID(id uint64) error {
	tx := repo.Store.Unscoped().Delete(&models.User{ID: id})
	return tx.Error
}
