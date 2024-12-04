package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define test models
type TestModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

type AnotherModel struct {
	ID    uint   `gorm:"primaryKey"`
	Value string `gorm:"size:255"`
}

func setupTestDB(t *testing.T) *gorm.DB {
	// Create an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err, "Failed to connect to the in-memory SQLite database")
	return db
}

func TestMigrate(t *testing.T) {
	db := setupTestDB(t)

	tables := []interface{}{&TestModel{}, &AnotherModel{}}
	err := Migrate(db, tables)

	assert.NoError(t, err, "Migrate should not return an error")
	for _, table := range tables {
		assert.True(t, db.Migrator().HasTable(table), "Table should be created after migration")
	}
}

func TestDrop(t *testing.T) {
	db := setupTestDB(t)

	// First migrate the tables
	tables := []interface{}{&TestModel{}, &AnotherModel{}}
	err := Migrate(db, tables)
	assert.NoError(t, err, "Migrate should not return an error")

	// Then drop the tables
	err = Drop(db, tables)
	assert.NoError(t, err, "Drop should not return an error")
	for _, table := range tables {
		assert.False(t, db.Migrator().HasTable(table), "Table should be dropped after calling Drop")
	}
}
