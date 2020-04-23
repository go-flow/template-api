package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/go-flow/flow"
	"github.com/go-flow/migrator"
	"github.com/go-flow/template-api/config"

	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// application data store
var store Store

// Init loads models for given application
func Init(app *flow.App) {
	cfg := app.AppConfig.(config.AppConfig)
	dbCfg, ok := cfg.DBConnections[app.Env]
	if !ok {
		app.Logger.Fatal(fmt.Sprintf(" Database connection configuration `%s` environment not provided", app.Env))
	}

	if dbCfg.DbDialect == "" || dbCfg.DbConnection == "" {
		// we consider that user does not want db connection
		return
	}

	// open DB connection
	db, err := gorm.Open(dbCfg.DbDialect, dbCfg.DbConnection)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}

	// ping DB
	if err = db.DB().Ping(); err != nil {
		app.Logger.Fatal(err.Error())
	}

	// SetMaxIdleConns sets maximum number of connections in the idle connection pool
	maxConn := dbCfg.DbMaxIdleConns
	db.DB().SetMaxIdleConns(maxConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database
	maxConn = dbCfg.DbMaxOpenConns
	db.DB().SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxConn = dbCfg.DbConnMaxLifetime
	duration := time.Minute * time.Duration(maxConn)
	db.DB().SetConnMaxLifetime(duration)

	// Enable Logger, show detailed log
	db.LogMode(dbCfg.DbLogging)

	// assign db object to application store
	store = db

	// register db store object to DI
	app.Register(store)

	// prepare DB migrations
	migrationsPath := cfg.DbMigrationsPath
	autoMigrate := cfg.DbMigrationsAutorun
	if migrationsPath == "" || !autoMigrate {
		return
	}

	app.Logger.Info("Start application migrations...")
	// execute migrations
	fm := migrator.NewFileMigrator(migrationsPath, dbCfg.DbDialect, db.DB())
	if err = fm.Up(); err != nil {
		app.Logger.Fatal(err.Error())
	}
	app.Logger.Info("End application migrations.")
}

// Close closes Database connection
//
// it is suitable for defer models.CloseDB()
func Close() error {
	if store != nil {
		return store.DB().Close()
	}
	return nil
}

// Connection gets database connection
func Connection() Store {
	return store
}
