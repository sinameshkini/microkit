package genericrepo

import (
	"context"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"gorm.io/gorm"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SomeModel struct {
	ID   int
	Name string
}

var conf = &database.Config{
	Host:     "localhost", // Update with your DB host
	Port:     "5432",      // Default PostgreSQL port
	Username: "admin",     // Your DB username
	Password: "admin",     // Your DB password
	Schema:   "test_db",   // Your test schema (ensure it's created in PostgreSQL)
	Debug:    true,
}

func setUp(db *gorm.DB) error {
	// Migrate models to ensure they exist before each test
	return database.Migrate(db, []interface{}{&SomeModel{}})
}

func tearDown(db *gorm.DB) {
	// Drop the tables after the test
	if err := database.Drop(db, []interface{}{&SomeModel{}}); err != nil {
		log.Fatal(err)
	}
	return
}

func TestRepository_Add(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Begin a transaction to isolate test data
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	// Act
	err = repo.Add(&SomeModel{ID: 1, Name: "Test"}, context.Background())

	// Assert
	assert.NoError(t, err)

	// Cleanup (rollback the transaction to avoid test pollution)
	tx.Rollback()
}

func TestRepository_GetById(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Insert a test record into the database
	err = db.Create(&SomeModel{ID: 1, Name: "Test"}).Error
	if err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	// Act
	entity, err := repo.GetById(1, context.Background())

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, entity)
	assert.Equal(t, 1, entity.ID)
	assert.Equal(t, "Test", entity.Name)

	// Cleanup
	db.Delete(&SomeModel{}, 1)
}

func TestRepository_Update(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Insert a test record into the database
	err = db.Create(&SomeModel{ID: 1, Name: "Test"}).Error
	if err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	// Act
	err = repo.Update(&SomeModel{ID: 1, Name: "Updated"}, context.Background())

	// Assert
	assert.NoError(t, err)

	// Cleanup
	db.Delete(&SomeModel{}, 1)
}

func TestRepository_Delete(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Insert a test record into the database
	err = db.Create(&SomeModel{ID: 1, Name: "Test"}).Error
	if err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	// Act
	err = repo.Delete(1, context.Background())

	// Assert
	assert.NoError(t, err)

	// Cleanup
	db.Delete(&SomeModel{}, 1)
}

func TestRepository_GetAll(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Insert test data into the database
	err = db.Create(&SomeModel{ID: 1, Name: "Test"}).Error
	if err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	req := &models.Request{
		GetPagination: true,
		PaginationRequest: models.PaginationRequest{
			Page:    1,
			PerPage: 10,
		},
	}

	// Act
	entities, meta, err := repo.GetAll(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, *entities, 1)
	assert.Equal(t, 1, (*entities)[0].ID)
	assert.Equal(t, "Test", (*entities)[0].Name)
	assert.NotNil(t, meta)

	// Cleanup
	db.Delete(&SomeModel{}, 1)
}

func TestRepository_AddAll(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Begin a transaction to isolate test data
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	// Act
	err = repo.AddAll(&[]SomeModel{{ID: 1, Name: "Test1"}, {ID: 2, Name: "Test2"}}, context.Background())

	// Assert
	assert.NoError(t, err)

	// Cleanup (rollback the transaction to avoid test pollution)
	tx.Rollback()
}

func TestRepository_SkipTake(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Insert test data into the database
	err = db.Create(&SomeModel{ID: 1, Name: "Test"}).Error
	if err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	// Act
	entities, err := repo.SkipTake(0, 10, context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Len(t, *entities, 1)
	assert.Equal(t, 1, (*entities)[0].ID)
	assert.Equal(t, "Test", (*entities)[0].Name)

	// Cleanup
	db.Delete(&SomeModel{}, 1)
}

func TestRepository_Count(t *testing.T) {
	// Arrange
	db, err := database.NewDBWithConf(conf)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	// Set up the database with migration
	if err := setUp(db); err != nil {
		t.Fatalf("failed to set up the database: %v", err)
	}
	defer tearDown(db) // Ensure the tables are dropped after the test

	repo := New[SomeModel](db)

	// Insert test data into the database
	err = db.Create(&SomeModel{ID: 1, Name: "Test"}).Error
	if err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	// Act
	count := repo.Count(context.Background())

	// Assert
	assert.Equal(t, int64(1), count)

	// Cleanup
	db.Delete(&SomeModel{}, 1)
}
