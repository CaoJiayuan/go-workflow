package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}) error {
	b, e := json.Marshal(data)
	if e != nil {
		return e
	}

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(b)
	return err
}
