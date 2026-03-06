package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/funciones"
	"log"
)

type Caballo struct {
	Nombre   string
	FechaNac string
	Sexo     string
	Peso     int
	Pelaje   string
	Madre    string
	Padre    string
}

func (u *Caballo) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Nombre, &u.FechaNac, &u.Sexo, &u.Peso, &u.Pelaje, &u.Padre, &u.Madre)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (u *Caballo) Load2(s, a, b, d, e, f string, c int) {

	u.Nombre = s
	u.FechaNac = a
	u.Sexo = b
	u.Peso = c
	u.Pelaje = d
	u.Padre = e
	u.Madre = f
}

func (p *Caballo) LoadDB() (int64, string) {

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

	res = db.Cliente.QueryRow(sqlStatement, p.Nombre, "", 3)
	err = res.Scan(&idd)

	if err == nil {
		return idd, m
	} else {
		return 0, "falla al obtener id: " + err.Error()
	}

}
