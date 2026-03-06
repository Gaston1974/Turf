package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/funciones"
	"log"
)

type Cuidador struct {
	Nombre   string
	Apellido string
	Ranking  int
}

func (u *Cuidador) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Nombre, &u.Apellido, &u.Ranking)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (u *Cuidador) Load2(s string, a string, b int) {

	u.Nombre = s
	u.Apellido = a
	u.Ranking = b

}

func (p *Cuidador) LoadDB() (int64, string) {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	var sqlStatement string
	var res *sql.Row

	var err error

	defer db.Cliente.Close()

	// Test the connection to the database
	if err := db.Cliente.Ping(); err != nil {
		log.Println(err)
		return 0, err.Error()
	} else {
		log.Println("\n Successfully Connected")
	}

	msg = "No se ha podido concretar el alta. "
	m := "Alta creada con exito"

	sqlStatement = "SELECT merge($1, $2, $3);"

	var idd int64

	res = db.Cliente.QueryRow(sqlStatement, p.Nombre, p.Apellido, 2)
	err = res.Scan(&idd)

	if err == nil {
		return idd, m
	} else {
		return 0, "falla al obtener id: " + err.Error()
	}

}
