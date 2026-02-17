package funciones

import (
	"fmt"
	"net/http"
)

func GetToken(headers http.Header) (string, string) {

	val := headers.Get("Authorization")

	if val == "" {
		msg := "\ntoken vacio"
		fmt.Printf("%s", msg)
		return "", msg
	}
	return val, ""
}
