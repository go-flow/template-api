package db

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Store is database store interface
type Store interface {
	Create(value interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	DB() *sql.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	RecordNotFound() bool
	Limit(limit interface{}) *gorm.DB
	Offset(offset interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	Not(query interface{}, args ...interface{}) *gorm.DB
	Debug() *gorm.DB
	LogMode(enable bool) *gorm.DB
	Related(value interface{}, foreignKeys ...string) *gorm.DB
	Unscoped() *gorm.DB
	Begin() *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
	Table(name string) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
}
