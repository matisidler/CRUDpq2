package product

import (
	"errors"
	"fmt"
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
	Migrate() error
	Create(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Update(*Model) error
	Delete(uint) error
}

//Servicio de producto
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

func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}

//Creamos un método para modelo, no influye en nada, solo hace que se vea más lindo al momento de hacer el Read.
func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s\n",
		m.ID, m.Name, m.Observaciones, m.Price, m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetById(i uint) (*Model, error) {
	return s.storage.GetById(i)
}

func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return errors.New("id not found")
	}
	return s.storage.Update(m)
}

func (s *Service) Delete(i uint) error {
	return s.storage.Delete(i)
}
