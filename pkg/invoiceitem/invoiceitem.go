package invoiceitem

import "time"

//Modelo de invoice item.
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
