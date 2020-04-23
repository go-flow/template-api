package repositories

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/models"
	"github.com/go-flow/template-api/pkg/paging"
)

// ValuesRepository defines set of available operations aroud Values model
type ValuesRepository interface {
	// ValuesRepository ensures interface implementation
	ValuesRepository() string

	// GetByID returns value model based on a provided id
	GetByID(id uint64) (*models.Value, error)

	// GetAll returns all Values model for given paging values
	GetAll(paginator *paging.Paginator) ([]*models.Value, error)

	// Create new Value object in database
	Create(value *models.Value) (*models.Value, error)

	// Update existing Value object in database
	Update(value *models.Value) (*models.Value, error)

	// Save creates or updates value model based on ID value
	Save(value *models.Value) (*models.Value, error)

	// Delete Value object from database
	Delete(value *models.Value) error
}

// NewValuesRepository creates new ValuesRepository implementation
func NewValuesRepository(app *flow.App) ValuesRepository {
	return &valuesRepository{}
}

// valuesRepository struct that implements ValuesRepository interface
type valuesRepository struct {
}

// ValuesRepository ensures interface implementation
func (svc *valuesRepository) ValuesRepository() string {
	return "valuesRepository"
}

// GetByID returns value model based on a provided id
func (svc *valuesRepository) GetByID(id uint64) (*models.Value, error) {
	return nil, nil
}

// GetAll returns all Values model for given paging values
func (svc *valuesRepository) GetAll(paginator *paging.Paginator) ([]*models.Value, error) {
	return nil, nil
}

// Create new Value object in database
func (svc *valuesRepository) Create(value *models.Value) (*models.Value, error) {
	return nil, nil
}

// Update existing Value object in database
func (svc *valuesRepository) Update(value *models.Value) (*models.Value, error) {
	return nil, nil
}

// Save creates or updates value model based on ID value
func (svc *valuesRepository) Save(value *models.Value) (*models.Value, error) {
	return nil, nil
}

// Delete Value object from database
func (svc *valuesRepository) Delete(value *models.Value) error {
	return nil
}
