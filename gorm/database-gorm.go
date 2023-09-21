package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
}

func main() {
	dsn := "host=localhost user=postgres password=root dbname=product port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	if err := db.AutoMigrate(&Product{}); err != nil {
		panic(err)
	}

	var product Product

	/** CREATE DATA
	INSERT INTO products (name, price) VALUES ('sari gandum', 20000)
	*/
	db.Create(&Product{Name: "sari gandum", Price: 20000})

	/** READ DATA
	SELECT * FROM "products" WHERE "products"."id" = 9 AND "products"."deleted_at" IS NULL ORDER BY "products"."id" LIMIT 1
	*/
	db.First(&product, 9) //find product with id

	/**
	SELECT * FROM "products" WHERE name = 'air' AND "products"."deleted_at" IS NULL ORDER BY "products"."id" LIMIT 1
	*/
	db.First(&product, "name = ?", "air") //find product with specific name
	fmt.Println(product)

	/** UPDATE DATA
	UPDATE "products" SET "name"='apalah',"updated_at"='2023-09-21 20:58:17.993' WHERE "products"."deleted_at" IS NULL AND "id" = 11
	*/
	db.Model(&product).Update("name", "apalah") //update one column

	/**
	UPDATE "products" SET "name"='air', "price" = 40000, "updated_at"='2023-09-21 20:58:17.993' WHERE "products"."deleted_at" IS NULL AND "id" = 11
	*/
	db.Model(&product).Updates(&Product{Name: "air", Price: 40000}) //update multiple column

	/** DELETE DATA
	UPDATE products SET deleted_at="2023-10-29 10:23" WHERE id = 11;
	*/
	db.Delete(&product, 11) //soft delete

	/**
	DELETE FROM products WHERE id = 8;
	*/
	db.Unscoped().Delete(&product, 8) //hard delete
}
