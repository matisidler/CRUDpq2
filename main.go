package main

import (
	"fmt"

	"github.com/matisidler/CRUDpqv2/pkg/invoice"
	"github.com/matisidler/CRUDpqv2/pkg/invoiceheader"
	"github.com/matisidler/CRUDpqv2/pkg/invoiceitem"
	"github.com/matisidler/CRUDpqv2/storage"
)

func main() {
	db := storage.NewPostgresDB()
	storageHeader := storage.PsqlNewInvoiceHeader(storage.Pool())
	storageItem := storage.PsqlNewInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(db, storageHeader, storageItem)

	serviceInvoice := invoice.NewService(storageInvoice)

	factura := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Matías",
		},
		Items: []*invoiceitem.Model{
			{ProductID: 4},
			{ProductID: 5},
			{ProductID: 6},
		},
	}
	err := serviceInvoice.Create(factura)
	if err != nil {
		fmt.Println(err)
	}
	/* db := storage.NewPostgresDB()
	storageItem := storage.PsqlNewInvoiceItem(db)
	storageHeader := storage.PsqlNewInvoiceHeader(db)
	err := storageHeader.Migrate()
	if err != nil {
		fmt.Println(err)
	}
	err = storageItem.Migrate()
	if err != nil {
		fmt.Println(err)
	} */

}

//Lo que quiere el programa es que en vez de pasarle un storage.PsqlInvoiceItem, le pase un invoiceitem.Storage
