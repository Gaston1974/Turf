package funciones

/*

SetCliente() ------------  Genera una conexion a la base de datos.
CreateUser() ------------  Crea usuario en la base de datos.
LogIn() -----------------  Valida que el usuario exista en la base de datos y la contraseña correspondiente.
ObtenerUsuario() --------  Obtiene un usuario de la base a partir del token.
ModifyPassword() --------  setea una nueva contraseña a partir del token.

*/

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"database/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

type Respuesta struct {
	Msg string
}

type Identificable interface {
	Load(*sql.Rows) string
	// Atr() []string
	// Get(string) string
}

type Acceso struct {
	Cliente *sql.DB
}

func (a *Acceso) SetCliente() (context.Context, int, string) {

	postgresURI := os.Getenv("DBURL")
	if postgresURI == "" {
		msg := "No URL variable is found on the environment"
		fmt.Printf("%s", msg)
		return nil, 0, msg
	}

	ctx := context.Background()
	var err error
	a.Cliente, err = sql.Open("postgres", postgresURI)

	if err != nil {
		msg := "No se ha logrado establecer conexion a la base de datos, intente mas tarde nuevamente. .."
		fmt.Printf("Cant connect to database : %s", err)
		return nil, 0, msg
	}

	return ctx, 1, ""

}

func LogIn(w http.ResponseWriter, r *http.Request) (int, string, int, string) {

	client := Acceso{}

	_, val0, msg := client.SetCliente()

	if val0 != 1 {
		return 400, msg, 0, ""
	}

	type parameters struct {
		NombreUs string
		Password string
	}

	b, err := io.ReadAll(r.Body)
	//fmt.Println(string(b))
	if err != nil {
		msg := "Falla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		return 400, msg, 0, ""
	}

	params := parameters{}

	err = json.Unmarshal(b, &params)
	if err != nil {
		msg := "\nFalla durante parseo de parametros del Request: "
		fmt.Printf("%s", msg)
		return 400, msg, 0, ""
	}

	//ctx := context.Background()

	defer client.Cliente.Close()

	hasher := sha256.New()
	hasher.Write([]byte(params.Password))
	passHash := hasher.Sum(nil)
	hashed := hex.EncodeToString(passHash)

	//fmt.Printf("\n username: %s \npass hashed : %s", params.NombreUs, hashed)

	row := client.Cliente.QueryRow("SELECT id, COALESCE(grado, ''), last_name FROM usuarios WHERE username = ? AND clave = ?;", params.NombreUs, hashed)

	if row == nil {
		msg := "\n contraseña o nombre de usuario invalido"
		fmt.Printf("%s row: %v", msg, row)
		return 401, msg, 0, ""
	}

	// ---------------
	var id int
	var grado string
	var nombre string

	err = row.Scan(&id, &grado, &nombre)
	if err != nil {
		msg := "\nFalla en la consulta a la base de datos"
		fmt.Printf("%s error: %s", msg, err.Error())
		return 400, msg, 0, ""
	}

	msg = "Usuario " + grado + " " + nombre + " logueado"
	fmt.Printf("\n\n %s id: %d", msg, id)
	return 1, msg, id, grado + " " + nombre
}

func ModifyPassword(w http.ResponseWriter, r *http.Request) (int, string) {

	client := Acceso{}

	_, val0, msg := client.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	type password struct {
		Id        string
		Password  string
		Password2 string
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "\nFalla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		return 0, msg
	}

	pass := password{}

	err = json.Unmarshal(b, &pass)
	if err != nil {
		msg := "falla durante parseo de parametros del Request: "
		fmt.Printf("%s", msg)
		return 0, msg
	}

	token, errorMsg := GetToken(r.Header)
	if errorMsg != "" {
		return 0, errorMsg
	}
	value, err := strconv.Atoi(token)
	if err != nil {
		msg := "Falla interna"
		fmt.Printf("\n %s en el casteo del token", msg)
		return 0, msg
	}

	//ctx := context.Background()

	defer client.Cliente.Close()

	hasher := sha256.New()
	//hasher.Write([]byte())
	hasher.Write([]byte(pass.Password))
	passHash := hasher.Sum(nil)
	hashed := hex.EncodeToString(passHash)

	sqlStatement := "UPDATE usuarios SET password = ? WHERE id = ?;"

	res, err := client.Cliente.Exec(sqlStatement, hashed, value)
	if err != nil {
		msg := "Falla al actualizar la contraseña"
		fmt.Printf("\n %s", msg)
		return 0, msg
	}

	count, err := res.RowsAffected()
	if err != nil {
		msg := "Falla al actualizar la contraseña"
		fmt.Printf("\n %s", msg)
		return 0, msg
	}
	fmt.Println(count)

	resp := Respuesta{}

	resp.Msg = "Contraseña modificada"
	ResponseWithJSON(w, 200, resp)

	return 1, ""
}
