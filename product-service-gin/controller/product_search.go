package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"pobsaeng.com/product-api/config"
	"pobsaeng.com/product-api/helpers"
)

func SearchProducts(c *gin.Context) {
	// Parse query parameters from the request body
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters", "details": err.Error()})
		return
	}

	// Build the base query and count query
	baseQuery := "SELECT * FROM tbl_product"
	countQuery := "SELECT COUNT(*) FROM tbl_product"

	// Build the full SQL query with filters based on the parsed parameters
	query := helpers.BuildSQLQuery(baseQuery, params)

	// Use the helper to fetch paginated data
	result, err := helpers.PaginateData(config.Db, query, countQuery, func(rows *sql.Rows) (interface{}, error) {
		var (
			id          sql.NullInt64
			code        sql.NullString
			name        sql.NullString
			description sql.NullString
			active      sql.NullBool
			price       sql.NullFloat64
			stock       sql.NullInt64
			weight      sql.NullFloat64
			brand       sql.NullString
			color       sql.NullString
			size        sql.NullString
			length      sql.NullFloat64
			width       sql.NullFloat64
			height      sql.NullFloat64
			image       sql.NullString
			categoryID  sql.NullInt64
			supplierID  sql.NullInt64
			createdBy   sql.NullString
			updatedBy   sql.NullString
			createdAt   sql.NullTime
			updatedAt   sql.NullTime
		)

		if err := rows.Scan(
			&id, &code, &name, &description, &active, &price, &stock, &weight,
			&brand, &color, &size, &length, &width, &height, &image, &categoryID,
			&supplierID, &createdBy, &updatedBy, &createdAt, &updatedAt,
		); err != nil {
			return nil, err
		}

		data := map[string]interface{}{
			"id":          nullableValue(id),
			"code":        nullableValue(code),
			"name":        nullableValue(name),
			"description": nullableValue(description),
			"active":      nullableValue(active),
			"price":       nullableValue(price),
			"stock":       nullableValue(stock),
			"weight":      nullableValue(weight),
			"brand":       nullableValue(brand),
			"color":       nullableValue(color),
			"size":        nullableValue(size),
			"length":      nullableValue(length),
			"width":       nullableValue(width),
			"height":      nullableValue(height),
			"image":       nullableValue(image),
			"category_id": nullableValue(categoryID),
			"supplier_id": nullableValue(supplierID),
			"created_by":  nullableValue(createdBy),
			"updated_by":  nullableValue(updatedBy),
			"created_at":  nullableValue(createdAt),
			"updated_at":  nullableValue(updatedAt),
		}

		return data, nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products", "details": err.Error()})
		return
	}

	// Send response as JSON
	c.JSON(http.StatusOK, result)
}


// nullableValue converts a sql.NullXXX type to a native Go type or nil if the value is NULL.
func nullableValue(value interface{}) interface{} {
	switch v := value.(type) {
	case sql.NullString:
		if v.Valid {
			return v.String
		}
	case sql.NullInt64:
		if v.Valid {
			return v.Int64
		}
	case sql.NullFloat64:
		if v.Valid {
			return v.Float64
		}
	case sql.NullBool:
		if v.Valid {
			return v.Bool
		}
	case sql.NullTime:
		if v.Valid {
			return v.Time
		}
	}
	return nil
}
