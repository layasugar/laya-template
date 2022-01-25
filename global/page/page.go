package page

import "math"

const (
	DefaultPageSize = 10
	MaxPageSize     = 100
	DefaultPage     = 1
)

type Pagination struct {
	CurrentPage uint `json:"current_page"` // 当前页码
	PerPage     uint `json:"per_page"`     // 当前页行数
	TotalPage   uint `json:"total_page"`   // 总页码
	Total       uint `json:"total"`        // 总行数
}

func GetPagination(page, pageSize uint, total uint) Pagination {
	return Pagination{
		Total:       total,
		CurrentPage: page,
		PerPage:     pageSize,
		TotalPage:   uint(math.Ceil(float64(total) / float64(pageSize))),
	}
}
