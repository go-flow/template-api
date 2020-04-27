package user

import (
	"fmt"

	"github.com/go-flow/template-api/db"
	"github.com/go-flow/template-api/domain/models"
	"github.com/go-flow/template-api/pkg/paging"

	"github.com/go-flow/flow"
)

// NewRepository creates new Repository implementation
func NewRepository(app *flow.App) *Repository {
	return &Repository{}
}

// Repository struct that implements Repository interface
type Repository struct {
	Store db.Store
}

// UserRepository ensures interface implementation
func (repo *Repository) UserRepository() string {
	return "Repository"
}

// GetByID returns user model based on a provided id
func (repo *Repository) GetByID(id uint64) (*models.User, error) {
	model := new(models.User)

	tx := repo.Store.Where("id = ?", id).First(model)

	return model, tx.Error
}

// GetByEmail returns user model based on a provided email
func (repo *Repository) GetByEmail(email string) (*models.User, error) {
	model := new(models.User)

	tx := repo.Store.Where("email = ?", email).First(model)

	return model, tx.Error
}

// GetAll returns all Users model for given paging users
func (repo *Repository) GetAll(paginator *paging.Paginator) ([]*models.User, error) {
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
func (repo *Repository) Create(user *models.User) error {
	tx := repo.Store.Create(user)
	return tx.Error
}

// Update existing User object in database
func (repo *Repository) Update(user *models.User) error {
	tx := repo.Store.Save(user)
	return tx.Error
}

// Save creates or updates user model based on ID user
func (repo *Repository) Save(user *models.User) error {
	if user.ID > 0 {
		return repo.Update(user)
	}
	return repo.Create(user)
}

// Delete user record from database
func (repo *Repository) Delete(user *models.User) error {
	return repo.DeleteByID(user.ID)
}

// DeleteByID user record from database
func (repo *Repository) DeleteByID(id uint64) error {
	tx := repo.Store.Unscoped().Delete(&models.User{ID: id})
	return tx.Error
}
