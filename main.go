package main

import (
	"log"

	"github.com/matisidler/CRUDpqv2/pkg/product"
	"github.com/matisidler/CRUDpqv2/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.ObtenerDB())
	serviceProduct := product.NewService(storageProduct)
	err := serviceProduct.Migrate()
	if err != nil {
		log.Fatalf("produc.Migrate: %v", err)
	}
}
