package repository

import (
	"database/sql"
	"log"
	"time"

	"pobsaeng.com/product-api/config"
	"pobsaeng.com/product-api/model"
)

func GetAllProducts() ([]model.Product, error) {
	db := config.Db
	query := `SELECT id, code, name, description, active, price, stock, weight, brand, color, size, length, width, height, image, category_id, supplier_id, created_by, updated_by, created_at, updated_at FROM tbl_product`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		var price, weight, length, width, height sql.NullString
		var description, brand, color, size, image, createdBy, updatedBy sql.NullString
		var categoryID, supplierID sql.NullInt64
		var active sql.NullBool
		var createdAt, updatedAt sql.NullTime

		err := rows.Scan(
			&product.ID, &product.Code, &product.Name, &description, &active,
			&price, &product.Stock, &weight, &brand, &color, &size,
			&length, &width, &height, &image, &categoryID, &supplierID,
			&createdBy, &updatedBy, &createdAt, &updatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}

		if price.Valid {
			product.Price.SetString(price.String)
		}
		if weight.Valid {
			product.Weight.SetString(weight.String)
		}
		if length.Valid {
			product.Length.SetString(length.String)
		}
		if width.Valid {
			product.Width.SetString(width.String)
		}
		if height.Valid {
			product.Height.SetString(height.String)
		}
		if description.Valid {
			product.Description = description.String
		}
		if brand.Valid {
			product.Brand = brand.String
		}
		if color.Valid {
			product.Color = color.String
		}
		if size.Valid {
			product.Size = size.String
		}
		if image.Valid {
			product.Image = image.String
		}
		if categoryID.Valid {
			product.CategoryID = uint64(categoryID.Int64)
		}
		if supplierID.Valid {
			product.SupplierID = uint64(supplierID.Int64)
		}
		if createdBy.Valid {
			product.CreatedBy = createdBy.String
		}
		if updatedBy.Valid {
			product.UpdatedBy = updatedBy.String
		}
		if createdAt.Valid {
			product.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			product.UpdatedAt = updatedAt.Time
		}
		if active.Valid {
			product.Active = active.Bool
		}

		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return products, nil
}

func GetProductByID(id uint64) (*model.Product, error) {
	db := config.Db
	query := `SELECT id, code, name, description, active, price, stock, weight, brand, color, size, length, width, height, image, category_id, supplier_id, created_by, updated_by, created_at, updated_at FROM tbl_product WHERE id = ?`
	var product model.Product
	var price, weight, length, width, height sql.NullString
	var description, brand, color, size, image, createdBy, updatedBy sql.NullString
	var categoryID, supplierID sql.NullInt64
	var active sql.NullBool
	var createdAt, updatedAt sql.NullTime

	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Code, &product.Name, &description, &active,
		&price, &product.Stock, &weight, &brand, &color, &size,
		&length, &width, &height, &image, &categoryID, &supplierID,
		&createdBy, &updatedBy, &createdAt, &updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Fatal(err)
	}

	if price.Valid {
		product.Price.SetString(price.String)
	}
	if weight.Valid {
		product.Weight.SetString(weight.String)
	}
	if length.Valid {
		product.Length.SetString(length.String)
	}
	if width.Valid {
		product.Width.SetString(width.String)
	}
	if height.Valid {
		product.Height.SetString(height.String)
	}
	if description.Valid {
		product.Description = description.String
	}
	if brand.Valid {
		product.Brand = brand.String
	}
	if color.Valid {
		product.Color = color.String
	}
	if size.Valid {
		product.Size = size.String
	}
	if image.Valid {
		product.Image = image.String
	}
	if categoryID.Valid {
		product.CategoryID = uint64(categoryID.Int64)
	}
	if supplierID.Valid {
		product.SupplierID = uint64(supplierID.Int64)
	}
	if createdBy.Valid {
		product.CreatedBy = createdBy.String
	}
	if updatedBy.Valid {
		product.UpdatedBy = updatedBy.String
	}
	if createdAt.Valid {
		product.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		product.UpdatedAt = updatedAt.Time
	}
	if active.Valid {
		product.Active = active.Bool
	}

	return &product, nil
}

func CreateProduct(product model.Product) error {
	db := config.Db
	query := `
	    INSERT INTO tbl_product 
	    (code, name, description, active, price, stock, weight, brand, color, size, length, width, height, image, category_id, supplier_id, created_by, updated_by) 
	    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
 
	_, err := db.Exec(query,
	    product.Code, product.Name, product.Description, product.Active, product.Price.String(),
	    product.Stock, product.Weight.String(), product.Brand, product.Color, product.Size,
	    product.Length.String(), product.Width.String(), product.Height.String(),
	    product.Image, product.CategoryID, product.SupplierID, product.CreatedBy,
	    product.UpdatedBy,
	)
 
	if err != nil {
	    log.Fatal(err)
	}
 
	return nil
 }
 

func UpdateProduct(id uint64, updatedProduct model.Product) error {
	db := config.Db
	query := `
	    UPDATE tbl_product 
	    SET code = ?, name = ?, description = ?, active = ?, price = ?, stock = ?, weight = ?, brand = ?, color = ?, size = ?, length = ?, width = ?, height = ?, image = ?, category_id = ?, supplier_id = ?, updated_by = ?, updated_at = ?
	    WHERE id = ?`
	_, err := db.Exec(query,
	    updatedProduct.Code, updatedProduct.Name, updatedProduct.Description, updatedProduct.Active, updatedProduct.Price.String(),
	    updatedProduct.Stock, updatedProduct.Weight.String(), updatedProduct.Brand, updatedProduct.Color, updatedProduct.Size,
	    updatedProduct.Length.String(), updatedProduct.Width.String(), updatedProduct.Height.String(),
	    updatedProduct.Image, updatedProduct.CategoryID, updatedProduct.SupplierID, updatedProduct.UpdatedBy,
	    time.Now(), id,
	)
 
	if err != nil {
	    log.Fatal(err)
	}
 
	return nil
 }

func DeleteProduct(id uint64) error {
	db := config.Db
	_, err := db.Exec("DELETE FROM tbl_product WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
