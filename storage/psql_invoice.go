package storage

import (
	"database/sql"

	"github.com/matisidler/CRUDpqv2/pkg/invoiceheader"
	"github.com/matisidler/CRUDpqv2/pkg/invoiceitem"
)

type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItem   invoiceitem.Storage
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItem:   i,
	}
}
