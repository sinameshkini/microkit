package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Request struct {
	PaginationRequest
	GetPagination bool      `json:"get_pagination" query:"get_pagination"`
	Preloads      []string  `json:"preloads" query:"preloads"`
	Filters       []*Filter `json:"filters"`
	Sorts         []*Sort   `json:"sorts"`
}

type Filter struct {
	Field     string      `json:"field"`
	Operation Operation   `json:"operation"`
	Value     interface{} `json:"value"`
}

type Sort struct {
	Field string        `json:"field"`
	Order SortDirection `json:"order"`
}

func (r *Request) AddToQuery(query *gorm.DB) *gorm.DB {
	for _, p := range r.Preloads {
		query = query.Preload(p)
	}

	for _, filter := range r.Filters {
		var (
			op string
			v  string = "?"
		)

		switch filter.Operation {
		case EQUAL:
			op = "="
		case NOTEQUAL:
			op = "<>"
		case GRATER:
			op = ">"
		case LOWER:
			op = "<"
		case EQUALGRATER:
			op = ">="
		case EQUALLOWER:
			op = "<="
		case IN:
			op = "in"
			v = "(?)"
		case NOTIN:
			op = "not in"
			v = "(?)"
		case LIKE:
			op = "like"
			v = "%?%"
		}

		query = query.Where(fmt.Sprintf("? %s %s", op, v), filter.Field, filter.Value)
	}

	for _, sort := range r.Sorts {
		query = query.Order(fmt.Sprintf("%s %s", sort.Field, sort.Order))
	}

	return query
}
