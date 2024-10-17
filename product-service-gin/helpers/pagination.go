package helpers

import (
	"database/sql"
	"fmt"
)

// Constants for pagination
const perPage = 5

// PaginationResult holds the pagination response data
type PaginationResult struct {
	Data      interface{} `json:"data"`
	Total     int64       `json:"total"`
	PageCount int         `json:"page_count"`
	LastPage  int         `json:"last_page"`
}

// BuildSQLQuery constructs a SQL query string based on filters and pagination
func BuildSQLQuery(baseQuery string, params map[string]interface{}) string {
	if searchTerm, ok := params["searchTerm"].(string); ok && searchTerm != "" {
		baseQuery += fmt.Sprintf(" WHERE name LIKE '%%%s%%' OR description LIKE '%%%s%%'", searchTerm, searchTerm)
	}

	if sortOrder, ok := params["sortOrder"].(string); ok && sortOrder != "" {
		baseQuery += fmt.Sprintf(" ORDER BY price %s", sortOrder)
	}

	page, ok := params["page"].(float64)
	if !ok {
		page = 1
	}
	offset := int((page - 1) * perPage)
	baseQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, offset)

	return baseQuery
}

// GetTotalCount retrieves the total count of records for pagination
func GetTotalCount(db *sql.DB, countQuery string) (int64, error) {
	var total int64
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// PaginateData executes the SQL query and returns the paginated data
func PaginateData(db *sql.DB, query string, countQuery string, scanFunc func(*sql.Rows) (interface{}, error)) (PaginationResult, error) {
	// Get the total count of records
	total, err := GetTotalCount(db, countQuery)
	if err != nil {
		return PaginationResult{}, err
	}

	// Execute the paginated query
	rows, err := db.Query(query)
	if err != nil {
		return PaginationResult{}, err
	}
	defer rows.Close()

	// Scan the rows using the provided scan function
	var results []interface{}
	for rows.Next() {
		result, err := scanFunc(rows)
		if err != nil {
			return PaginationResult{}, err
		}
		results = append(results, result)
	}

	// Handle any potential errors during iteration
	if err = rows.Err(); err != nil {
		return PaginationResult{}, err
	}

	// Calculate last page
	lastPage := int((total + perPage - 1) / perPage)

	return PaginationResult{
		Data:      results,
		Total:     total,
		PageCount: len(results),
		LastPage:  lastPage,
	}, nil
}
