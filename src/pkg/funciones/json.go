package funciones

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v ", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(dat)

}
