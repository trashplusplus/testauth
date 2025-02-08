package product

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Id    int
	Title string
	Price int
}

func GetProducts(db *sql.DB) []Product {

	var products []Product

	sqlStatement := "select * from products order by price desc"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("GetProducts Error: %s", err)
		return nil
	}

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.Title, &product.Price)
		if err != nil {
			fmt.Println("GetProducts Error: %s", err)
			return nil
		}

		products = append(products, product)

	}

	return products
}
