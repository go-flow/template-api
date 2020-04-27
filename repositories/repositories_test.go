package repositories

import (
	"database/sql/driver"
	"time"
)

// AnyTime used to mock time object
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
