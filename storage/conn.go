//ACÁ NOS CONECTAMOS A LA BD

package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

//Creamos la conexión a la BD.
//Utilizamos el Patrón Singleton para que solo se ejecute una vez.

//Creamos dos variables que van a poder ser usadas por todos los archivos del paquete storage.
//Con once hacemos el patrón Singleton para que se ejecute una sola vez.
var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:password@localhost:5432/gocrud?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open DB %v", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("can't do ping: %v", err)
		}
		fmt.Println("Conectado a postgres.")
	})
	return db
}

func stringToNull(s string) sql.NullString {
	var nullString sql.NullString
	if s == "" {
		nullString.Valid = false
	} else {
		nullString.Valid = true
		nullString.String = s
	}
	return nullString
}
