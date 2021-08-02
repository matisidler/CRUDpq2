package storage

import (
	"database/sql"
	"fmt"
)

//Creamos una constante (como mi variable "q") para ejecutar las querys.
const (
	//CONSTAINT: por defecto se pone asi: nombreTabla_nombreColumna_primaryKey/foreignKey
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)

	) `
)

//PsqlProduct nos genera la variable db para interactuar con la base de datos.
type PsqlProduct struct {
	db *sql.DB
}

//NewPsqlProduct retorna un nuevo puntero de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

//Migrate crea la tabla Products en la base de datos
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de Producto ejecutada correctamente")
	return nil
}
