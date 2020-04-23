package config

import (
	"log"
	"os"
	"strconv"

	"github.com/go-flow/flow"
)

type dbConnection struct {
	DbDialect         string
	DbConnection      string
	DbMaxIdleConns    int
	DbMaxOpenConns    int
	DbConnMaxLifetime int
	DbLogging         bool
}

// AppConfig - application configuration
type AppConfig struct {
	DBConnections       map[string]dbConnection
	DbMigrationsPath    string
	DbMigrationsAutorun bool
}

// Load application configuration
func Load() flow.Options {
	return LoadWithVersion("v0.0.0")

}

// LoadWithVersion loads application configuration with build version
func LoadWithVersion(version string) flow.Options {
	cfg := AppConfig{
		DBConnections: map[string]dbConnection{
			"development": {
				DbDialect:         "mysql",
				DbConnection:      getEnv("DB_DEV_CONNECTION", "root:root@(localhost:3306)/flow_dev?multiStatements=true&readTimeout=1800s&charset=utf8mb4&parseTime=True&loc=Local"),
				DbMaxIdleConns:    10,
				DbMaxOpenConns:    100,
				DbConnMaxLifetime: 30, // minutes
				DbLogging:         true,
			},
			"test": {
				DbDialect:         "mysql",
				DbConnection:      getEnv("DB_TEST_CONNECTION", "root:root@(localhost:3306)/flow_dev?multiStatements=true&readTimeout=1800s&charset=utf8mb4&parseTime=True&loc=Local"),
				DbMaxIdleConns:    10,
				DbMaxOpenConns:    100,
				DbConnMaxLifetime: 30, // minutes
				DbLogging:         true,
			},
			"production": {
				DbDialect:         "mysql",
				DbConnection:      getEnv("DB_PROD_CONNECTION", "root:root@(localhost:3306)/flow_dev?multiStatements=true&readTimeout=1800s&charset=utf8mb4&parseTime=True&loc=Local"),
				DbMaxIdleConns:    10,
				DbMaxOpenConns:    100,
				DbConnMaxLifetime: 30, // minutes
				DbLogging:         false,
			},
		},
		DbMigrationsPath:    "./migrations",
		DbMigrationsAutorun: true,
	}

	// get application options
	opts := flow.NewOptions()
	opts.Env = getEnv("ENV", "development")
	opts.Addr = getEnv("ADDR", opts.Addr)
	opts.LogLevel = getEnv("LOG_LEVEL", "debug")
	opts.Version = version
	opts.HandleMethodNotAllowed = true
	opts.RedirectTrailingSlash = false
	opts.RequestLoggerIgnore = []string{"/health/"}

	// set appConfig to options
	opts.AppConfig = cfg

	return opts
}

// getEnv returns value for given key from environment
// if key is not present in environment it returns defaultValue
func getEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) > 0 {
		return v
	}
	return defaultValue
}

// getEnvInt returns integer value for given key from environment
// if key is not present in environment it returns defaultValue
// if key cannot be parsed to integer function will panic
func getEnvInt(key string, defaultValue int) int {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultValue
	}

	valInteger, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf(" variable `%s` cannot be parsed to INTEGER", key)

	}

	return valInteger
}

// getEnvFloat returns float value for given key from environment
// if key is not present in environment it returns defaultValue
// if key cannot be parsed to float function will panic
func getEnvFloat(key string, defaultValue float32) float32 {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultValue
	}

	valFloat, err := strconv.ParseFloat(v, 32)
	if err != nil {
		log.Fatalf(" variable `%s` cannot be parsed to FLOAT", key)

	}

	return float32(valFloat)
}

// mustGetEnv returns value for given key from environment
// if key is not present in environment function will panic
func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		log.Fatalf(" variable `%s` is not present in ENVIRONMENT", key)
	}
	return v
}
