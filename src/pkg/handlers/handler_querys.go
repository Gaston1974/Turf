package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	funciones "hello/src/pkg/funciones"
)

func HandlerCombos(w http.ResponseWriter, r *http.Request) {

	db := funciones.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		funciones.ResponseWithJSON(w, 400, msg)
	}

	defer db.Cliente.Close()

	var err error
	var sqlStatement string
	var rows *sql.Rows

	type entidad1 struct {
		Id          string
		Descripcion string
	}

	type entidad2 struct {
		Nombre string
	}

	tabla := entidad1{}
	object := []entidad1{}

	tabla2 := entidad2{}
	object2 := []entidad2{}

	path := r.URL.Path

	vector := strings.Split(path, "/")
	value := vector[3]

	if value == "preventores" {

		sqlStatement = " SELECT nombre FROM " + value +
			" ORDER BY nombre;"

	} else {

		sqlStatement = " SELECT id, nombre FROM " + value +
			" WHERE is_active = 1 " +
			" ORDER BY id;"

	}

	rows, err = db.Cliente.Query(sqlStatement)

	if err != nil {

		funciones.ResponseWithJSON(w, 400, "error en la consulta a la base de datos")
		return

	} else if value == "preventores" {
		for rows.Next() {
			rows.Scan(&tabla2.Nombre)

			object2 = append(object2, tabla2)
		}

		funciones.ResponseWithJSON(w, 200, object2)

	} else {

		for rows.Next() {
			rows.Scan(&tabla.Id, &tabla.Descripcion)

			object = append(object, tabla)
		}

		funciones.ResponseWithJSON(w, 200, object)

	}

}
