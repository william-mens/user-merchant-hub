package helpers

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	CurrentPage  int `json:"currentPage"`
	PreviousPage int `json:"previousPage"`
	NextPage     int `json:"nextPage"`
	TotalRecords int `json:"totalRecords"`
	FromPage     int `json:"fromPage"`
	ToPage       int `json:"toPage"`
	TotalPages   int `json:"totalPages"`
	PerPage      int `json:"perPage"`
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	offset := (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	}
}

// Pagination metadata structure
func GetPaginationMetadata(page, totalRecords, limit int) Pagination {
	var (
		currentPage  = page
		previousPage = 0
		nextPage     = 0
		fromPage     = 1
		toPage       = 0
	)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if page > 1 {
		previousPage = page - 1
	}
	if page < totalPages {
		nextPage = page + 1
	}

	toPage = page * limit
	fromPage = (page-1)*limit + 1
	if toPage > totalRecords {
		toPage = totalRecords
	}

	return Pagination{
		CurrentPage:  currentPage,
		PreviousPage: previousPage,
		NextPage:     nextPage,
		TotalRecords: totalRecords,
		FromPage:     fromPage,
		ToPage:       toPage,
		TotalPages:   totalPages,
		PerPage:      limit,
	}
}
