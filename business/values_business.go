package business

import (
	"github.com/go-flow/flow"
	"github.com/go-flow/template-api/models"
	"github.com/go-flow/template-api/pkg/paging"
)

// ValuesBusiness defines set of business rules related to Values model
type ValuesBusiness interface {
	// ValuesBusiness ensures interface implementation
	ValuesBusiness() string

	// GetByID returns value model based on a provided id
	GetByID(id uint64) (*models.Value, error)

	// GetAll returns all Values model for given paging values
	GetAll(paginator *paging.Paginator) ([]*models.Value, error)

	// Create new Value object in database
	Create(name string, description string) (*models.Value, error)

	// Update existing Value object in database
	Update(id uint64, name string, description string) (*models.Value, error)

	// Delete Value object from database
	Delete(id uint64) error
}

// NewValuesBusiness creates new values business rules implementation instance
func NewValuesBusiness(app *flow.App) ValuesBusiness {
	return &valuesBusiness{}
}

// valuesBusiness struct that implements ValuesBusiness interface
type valuesBusiness struct {
}

// ValuesBusiness ensures interface implementation
func (bl *valuesBusiness) ValuesBusiness() string {
	return "valuesBusiness"
}

// GetByID returns value model based on a provided id
func (bl *valuesBusiness) GetByID(id uint64) (*models.Value, error) {
	return nil, nil
}

// GetAll returns all Values model for given paging values
func (bl *valuesBusiness) GetAll(paginator *paging.Paginator) ([]*models.Value, error) {
	return nil, nil
}

// Create new Value object in database
func (bl *valuesBusiness) Create(name string, description string) (*models.Value, error) {
	return nil, nil
}

// Update existing Value object in database
func (bl *valuesBusiness) Update(id uint64, name string, description string) (*models.Value, error) {
	return nil, nil
}

// Delete Value object from database
func (bl *valuesBusiness) Delete(id uint64) error {
	return nil
}
