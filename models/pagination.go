package models

import (
	"gorm.io/gorm"
	"math"
)

type PaginationRequest struct {
	Page    int64 `json:"page" query:"page"`
	PerPage int64 `json:"per_page" query:"per_page"`
}

func (r *PaginationRequest) Normalize() {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PerPage <= 0 {
		r.PerPage = 10
	}
}

func (r *PaginationRequest) ToQuery(query *gorm.DB) *gorm.DB {
	r.Normalize()
	return query.Limit(int(r.PerPage)).Offset(int(r.PerPage * (r.Page - 1)))
}

func GetCount(db *gorm.DB) (total int64, err error) {
	if db.Count(&total).Error != nil {
		return 0, err
	}

	return total, nil
}

type PaginationResponse struct {
	Total       int64 `json:"total"`
	TotalPages  int64 `json:"total_pages"`
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}

func MakePaginationResponse(total, page, perPage int64) *PaginationResponse {
	totalPages := int64(math.Ceil(float64(total) / float64(perPage)))
	hasNext := totalPages > page
	hasPrevious := page > 1

	return &PaginationResponse{
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: page,
		PerPage:     perPage,
		HasNext:     hasNext,
		HasPrevious: hasPrevious,
	}
}
