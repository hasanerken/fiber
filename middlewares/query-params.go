package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"strconv"
	"strings"
	"xhantos/common"
)

func QueryParamsMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		queryParams := common.QueryParams{
			SortBy:  ctx.Query("sortBy", ""),
			IsDesc:  true,
			Filters: make(map[string]interface{}),
		}

		// Convert page and perPage parameters to integers
		page, err := strconv.Atoi(ctx.Query("page", "1"))
		if err != nil || page <= 0 {
			page = 1
		}
		queryParams.Page = page

		perPage, err := strconv.Atoi(ctx.Query("perPage", "10"))
		if err != nil || perPage <= 0 {
			perPage = 10 // Default value for number of items per page
		}
		queryParams.PerPage = perPage

		// Check if the sortBy parameter contains the "-" prefix, indicating descending order
		if strings.HasPrefix(queryParams.SortBy, "-") {
			queryParams.IsDesc = true
			// Remove the "-" prefix from the sortBy parameter to get the actual field name
			queryParams.SortBy = queryParams.SortBy[1:]
		}

		// Check if the isDesc parameter is set to "false" in the query string
		isDescStr := ctx.Query("isDesc", "")
		if isDescStr != "" {
			isDesc, err := strconv.ParseBool(isDescStr)
			if err == nil {
				queryParams.IsDesc = isDesc
			}
		}

		// Manually parse the query string to extract filters
		queryString := string(ctx.Request().URI().QueryString())
		queryValues, err := url.ParseQuery(queryString)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse query parameters",
			})
		}

		// Add filters from query parameters to the map
		for key, values := range queryValues {
			if key != "sortBy" && key != "page" && key != "perPage" && key != "isDesc" {
				// For any query parameter other than sortBy, page, perPage, and isDesc, assume it's a filter
				if len(values) > 0 {
					queryParams.Filters[key] = values[0]
				}
			}
		}

		// Store queryParams in ctx.Locals
		ctx.Locals("queryParams", queryParams)

		// Continue to the next middleware/route handler
		return ctx.Next()
	}
}
