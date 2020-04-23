package services

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/models"
	"github.com/go-flow/template-api/pkg/paging"
	"github.com/go-flow/template-api/repositories"
)

// ValuesService defines set of available operations aroud Values model
type ValuesService interface {
	// ValuesService ensures interface implementation
	ValuesService() string

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

// NewValuesService creates new ValuesService implementation
func NewValuesService(app *flow.App) ValuesService {
	return &valuesService{}
}

// valuesService struct that implements ValuesService interface
type valuesService struct {
	// ValuesRepository implementation injected by dependency injection
	ValuesRepository repositories.ValuesRepository
}

// ValuesService ensures interface implementation
func (svc *valuesService) ValuesService() string {
	return "valuesService"
}

// GetByID returns value model based on a provided id
func (svc *valuesService) GetByID(id uint64) (*models.Value, error) {
	return svc.ValuesRepository.GetByID(id)
}

// GetAll returns all Values model for given paging values
func (svc *valuesService) GetAll(paginator *paging.Paginator) ([]*models.Value, error) {
	return svc.ValuesRepository.GetAll(paginator)
}

// Create new Value object in database
func (svc *valuesService) Create(value *models.Value) (*models.Value, error) {
	return svc.ValuesRepository.Create(value)
}

// Update existing Value object in database
func (svc *valuesService) Update(value *models.Value) (*models.Value, error) {
	return svc.ValuesRepository.Update(value)
}

// Save creates or updates value model based on ID value
func (svc *valuesService) Save(value *models.Value) (*models.Value, error) {
	if value.ID > 0 {
		return svc.Update(value)
	}

	return svc.Create(value)
}

// Delete Value object from database
func (svc *valuesService) Delete(value *models.Value) error {
	return svc.ValuesRepository.Delete(value)
}
