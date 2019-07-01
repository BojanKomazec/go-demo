package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BojanKomazec/go-demo/internal/pkg/dbclient"
)

type (
	// DbConfig represents DataBase configuration that shall be obtained from the runtime environment
	DbConfig struct {
		connParams dbclient.ConnParams
	}

	// Config represents application configuration that shall be obtained from the runtime environment
	Config struct {
		DB        DbConfig
		OutputDir string
	}

	keyValueReader = func(key string) (value string, exists bool)
)

// ConnParams returns connParams member of DbConfig
func (dbConfig DbConfig) ConnParams() dbclient.ConnParams {
	return dbConfig.connParams
}

// New returns a new Config struct
func New() (*Config, error) {
	var kvr = func(key string) (value string, exists bool) {
		return os.LookupEnv(key)
	}

	return newImpl(kvr)
}

func newImpl(kvr keyValueReader) (*Config, error) {
	dbHost, err := getStringValue("DB_HOST", kvr)
	if err != nil {
		return nil, err
	}

	dbPort, err := getPortNumber("DB_PORT", kvr)
	if err != nil {
		return nil, err
	}

	dbName, err := getStringValue("DB_NAME", kvr)
	if err != nil {
		return nil, err
	}

	dbUser, err := getStringValue("DB_USER", kvr)
	if err != nil {
		return nil, err
	}

	dbPass, err := getStringValue("DB_PASSWORD", kvr)
	if err != nil {
		return nil, err
	}

	outputDir, err := getStringValue("OUTPUT_DIR", kvr)
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: DbConfig{
			connParams: dbclient.NewConnParams(
				dbHost,
				dbPort,
				dbName,
				dbUser,
				dbPass),
		},
		OutputDir: outputDir,
	}, nil
}

// getStringValue returns a non-empty string read from env variable or an error
func getStringValue(key string, kvr keyValueReader) (string, error) {
	host, err := getEnv(key, kvr)
	if err != nil {
		return "", err
	}

	if host == "" {
		return "", fmt.Errorf("%s value is empty string", key)
	}

	return host, nil
}

func getPortNumber(key string, kvr keyValueReader) (int, error) {
	port, err := getEnvAsInt(key, kvr)
	if err != nil {
		return 0, err
	}

	const minPort, maxPort = 0, 65535
	if port < minPort || port > maxPort {
		return 0, fmt.Errorf("Port is out of valid range")
	}

	return port, nil
}

// Helper function to read an environment variable
func getEnv(key string, r keyValueReader) (string, error) {
	value, found := r(key)

	if !found {
		return "", fmt.Errorf("Environment variable \"%s\" does not exist", key)
	}

	return value, nil
}

// Helper function to read an environment variable into integer
func getEnvAsInt(name string, r keyValueReader) (int, error) {
	valueStr, err := getEnv(name, r)

	if err != nil {
		return 0, err
	}

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value, nil
	}

	return 0, fmt.Errorf("Error parsing \"%s\" as Int", valueStr)
}
