package product

import (
	"time"
)

//Modelo de producto
type Model struct {
	ID            uint
	Name          string
	Observaciones string
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//Creamos un slice de *Model para la función GetAll, que nos va a devolver un slice de todos los modelos que hay en la tabla, ya que es la función que lee.
type Models []*Model

//La interfaz storage es como mi archivo crud.go en el CRUD anterior. Me permite ejecutar las funciones para crear, borrar, actualizar y leer.
type Storage interface {
	//Create recibe como parametro un Modelo de Producto/Invoice/InvoiceHeader para crear uno de estos. Devuelve un error en caso de que exista
	Migrate() error
	/* Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error */
}

//Servicio de producto
type Service struct {
	storage Storage
}

//Retorna un puntero de Service
func NewService(s Storage) *Service {
	return &Service{}
}

//Migrate se usa para migrar producto. Es decir, crear la tabla producto.
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
