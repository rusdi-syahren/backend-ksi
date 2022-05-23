package shared

// AllowedSortFields is allowed field name for sorting
var AllowedSortFields = []string{
	"name",
	"id",
	"created",
	"distance",
	"rating",
	"lastModified",
	"store_id",
}

// Parameters data structure
type Parameters struct {
	StrPage        string
	Page           int
	StrLimit       string
	Limit          int
	Offset         int
	IsDeleteString string
	IsDelete       bool
	Sort           string
	OrderBy        string
	DateFrom       string
	DateTo         string
	CreatorID      string
	Query          string
}
