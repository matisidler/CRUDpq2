package main

import (
	"fmt"

	"github.com/matisidler/CRUDpqv2/pkg/product"
	"github.com/matisidler/CRUDpqv2/storage"
)

func main() {
	db := storage.NewPostgresDB()
	productStorage := storage.NewPsqlProduct(db)
	productService := product.NewService(productStorage)
	err := productService.Delete(9)
	if err != nil {
		fmt.Println(err)
	}

	cs, err := productService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cs)
}
