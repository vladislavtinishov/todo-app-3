package utils

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
}

// Apply adds pagination to query. Add this at the end to correctly calculate the total number of rows.
func (p *Pagination) Apply(query *gorm.DB) (*gorm.DB, error) {
	if err := query.Count(&p.Total).Error; err != nil {
		return nil, err
	}

	if p.Page < 1 || p.Limit < 1 {
		p.Page = 1
		p.Limit = 15
	}

	return query.Limit(p.Limit).Offset((p.Page - 1) * p.Limit), nil
}
