package invoiceitem

import (
	"database/sql"
	"time"
)

//Modelo de invoice item.
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, uint, []*Model) error
}

//Servicio de InvoiceItem
type Service struct {
	storage Storage
}

//Retorna un puntero de Service
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate se usa para migrar producto. Es decir, crear la tabla producto.
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

func (s *Service) CreateTx(t *sql.Tx, i uint, m []*Model) error {
	return s.storage.CreateTx(t, i, m)
}
