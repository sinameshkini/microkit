package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestPaginationRequest_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		request  PaginationRequest
		expected PaginationRequest
	}{
		{
			name:     "Valid values",
			request:  PaginationRequest{Page: 2, PerPage: 15},
			expected: PaginationRequest{Page: 2, PerPage: 15},
		},
		{
			name:     "Zero page and per_page",
			request:  PaginationRequest{Page: 0, PerPage: 0},
			expected: PaginationRequest{Page: 1, PerPage: 10},
		},
		{
			name:     "Negative page and per_page",
			request:  PaginationRequest{Page: -3, PerPage: -20},
			expected: PaginationRequest{Page: 1, PerPage: 10},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.request.Normalize()
			assert.Equal(t, test.expected, test.request)
		})
	}
}

func TestGetCount(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&testModel{})
	db.Create(&testModel{Name: "A"})
	db.Create(&testModel{Name: "B"})

	total, err := GetCount(db.Model(&testModel{}))
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
}

func TestMakePaginationResponse(t *testing.T) {
	tests := []struct {
		name          string
		total         int64
		page          int64
		perPage       int64
		expectedPages int64
		hasNext       bool
		hasPrevious   bool
	}{
		{
			name:          "Middle page",
			total:         50,
			page:          2,
			perPage:       10,
			expectedPages: 5,
			hasNext:       true,
			hasPrevious:   true,
		},
		{
			name:          "First page",
			total:         10,
			page:          1,
			perPage:       10,
			expectedPages: 1,
			hasNext:       false,
			hasPrevious:   false,
		},
		{
			name:          "Last page",
			total:         25,
			page:          3,
			perPage:       10,
			expectedPages: 3,
			hasNext:       false,
			hasPrevious:   true,
		},
		{
			name:          "Empty data",
			total:         0,
			page:          1,
			perPage:       10,
			expectedPages: 0,
			hasNext:       false,
			hasPrevious:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := MakePaginationResponse(test.total, test.page, test.perPage)
			assert.Equal(t, test.total, response.Total)
			assert.Equal(t, test.expectedPages, response.TotalPages)
			assert.Equal(t, test.page, response.CurrentPage)
			assert.Equal(t, test.perPage, response.PerPage)
			assert.Equal(t, test.hasNext, response.HasNext)
			assert.Equal(t, test.hasPrevious, response.HasPrevious)
		})
	}
}

type testModel struct {
	ID   uint
	Name string
}
