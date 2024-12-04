package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_dsn(t *testing.T) {
	conf := &Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "test_user",
		Password: "test_password",
		Schema:   "test_db",
		Debug:    true,
	}

	expectedDsn := "host=localhost user=test_user password=test_password dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Tehran"
	assert.Equal(t, expectedDsn, conf.dsn(), "Generated DSN should match the expected value")
}

func TestNewDBWithConf(t *testing.T) {
	conf := &Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "admin",
		Password: "admin",
		Schema:   "test_db",
		Debug:    true,
	}

	db, err := NewDBWithConf(conf)
	assert.NoError(t, err, "NewDBWithConf should not return an error for a valid config")
	assert.NotNil(t, db, "NewDBWithConf should return a non-nil *gorm.DB instance")
}

func TestNewDBWithDsn(t *testing.T) {
	dsn := "host=localhost user=admin password=admin dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := NewDBWithDsn(dsn, true)

	assert.NoError(t, err, "NewDBWithDsn should not return an error for a valid DSN")
	assert.NotNil(t, db, "NewDBWithDsn should return a non-nil *gorm.DB instance")
}

func TestGetLogger(t *testing.T) {
	debugLogger := getLogger(true)
	errorLogger := getLogger(false)

	assert.NotNil(t, debugLogger, "Logger should not be nil when debug is true")
	assert.NotNil(t, errorLogger, "Logger should not be nil when debug is false")
}
