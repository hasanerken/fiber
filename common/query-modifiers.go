package common

import (
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type QueryParams struct {
	SortBy  string
	IsDesc  bool
	Page    int
	PerPage int
	Filters map[string]interface{}
}

func BuildQueryModifiers(queryParams QueryParams) []qm.QueryMod {
	modifiers := []qm.QueryMod{
		qm.Where("1 = 1"), // Placeholder to add more conditions as needed
	}

	// Apply filters
	for field, value := range queryParams.Filters {
		modifiers = append(modifiers, qm.Where(fmt.Sprintf("%s = ?", field), value))
	}

	// Apply sorting based on SortBy and IsDesc
	if queryParams.SortBy != "" {
		if queryParams.IsDesc {
			modifiers = append(modifiers, qm.OrderBy(fmt.Sprintf("%s DESC", queryParams.SortBy)))
		} else {
			modifiers = append(modifiers, qm.OrderBy(queryParams.SortBy))
		}
	}

	// Apply pagination based on Page and PerPage
	page := queryParams.Page
	perPage := queryParams.PerPage
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10 // Default value for number of items per page
	}

	offset := (page - 1) * perPage
	modifiers = append(modifiers, qm.Limit(perPage), qm.Offset(offset))

	return modifiers

}
