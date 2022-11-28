package pagination

const (
	DefaultPageSize = 10
	MaxPageSize     = 100
	DefaultPage     = 1
)

type Pagination struct {
	Count  uint64 `json:"count"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

func GetPagination(count, offset, limit uint64) Pagination {
	return Pagination{
		Count:  count,
		Offset: offset,
		Limit:  limit,
	}
}
