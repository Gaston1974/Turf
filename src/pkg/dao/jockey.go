package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/funciones"
	"log"
)

type Jockey struct {
	Nombre   string
	Apellido string
	Ranking  int
}

func (u *Jockey) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Nombre, &u.Apellido, &u.Ranking)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (u *Jockey) Load2(s string, a string, b int) {

	u.Nombre = s
	u.Apellido = a
	u.Ranking = b

}

func (p *Jockey) LoadDB() (int, string) {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	var sqlStatement string
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

	sqlStatement = "INSERT INTO jockey  " +
		" (nombre, apellido, ranking) VALUES ($1, $2, $3);"

	res, err = db.Cliente.Exec(sqlStatement, p.Nombre, p.Apellido, p.Ranking)

	if err != nil {

		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0, msg
	} else {
		i, _ := res.RowsAffected()
		fmt.Printf("\n count: %d", int(i))
		return 1, m
	}

}
