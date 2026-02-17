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

func (p *Caballo) LoadDB() (int, string) {

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

	sqlStatement = "INSERT INTO caballo  " +
		" (nombre, fecha_nac, sexo, pelaje, padre, madre, peso) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	res, err = db.Cliente.Exec(sqlStatement, p.Nombre, p.FechaNac, p.Sexo, p.Pelaje, p.Padre, p.Madre, p.Padre)

	if err != nil {

		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0, msg
	} else {
		i, _ := res.RowsAffected()
		fmt.Printf("\n count: %d", int(i))
		return 1, m
	}

}
