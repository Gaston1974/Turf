package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/funciones"
	"log"
)

type Carrera struct {
	Fecha       string
	Nombre      string
	Pista       string
	Distancia   int64
	Descripcion string
	Detalle     []CarreraDet
}

func (u *Carrera) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Fecha, &u.Nombre, &u.Pista, &u.Distancia)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (u *Carrera) Load2(s string, a string, b string, c int64) {

	u.Fecha = s
	u.Nombre = a
	u.Pista = b
	u.Distancia = c

}

func (p *Carrera) LoadDB() (int, string) {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	var sqlStatement, sqlStatement2, sqlStatement3 string
	var res sql.Result

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

	sqlStatement = "INSERT INTO carrera  " +
		" (fecha, nombre, pista, distancia, descripcion) VALUES ($1, $2, $3, $4, $5);"

	sqlStatement2 = "INSERT INTO carrera_detalle  " +
		" (nombre, competidor, jockey, cuidador, handicap) VALUES ($1, $2, $3, $4, $5);"

	sqlStatement3 = "DELETE RFOM carrera WHERE nombre = $1; "

	res, err = db.Cliente.Exec(sqlStatement, p.Fecha, p.Nombre, p.Pista, p.Distancia, p.Descripcion)

	if err != nil {

		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0, msg

	} else {
		i, _ := res.RowsAffected()
		fmt.Printf("\n count carrera: %d", int(i))

		for _, v := range p.Detalle {

			res, err = db.Cliente.Exec(sqlStatement2, v.Nombre, v.Competidor, v.Jockey, v.Cuidador, v.Handicap)

			if err != nil {

				fmt.Printf("\n %s error: %s", msg, err.Error())
				res, err = db.Cliente.Exec(sqlStatement3, p.Nombre)
				return 0, msg

			} else {
				i, _ := res.RowsAffected()
				fmt.Printf("\n count carrera detalle: %d", int(i))

			}

		}

	}

	return 1, m

}
