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

func NewPostgresDB() {
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
}

//Tenemos que crear una función que nos devuelva la variable DB. Esto es asi porque con el patrón Singleton se debe hacer esto
//Para que los demás paquetes puedan usar la variable db.
func ObtenerDB() *sql.DB {
	return db
}
