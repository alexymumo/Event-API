package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ERROR(w http.ResponseWriter, statuscode int, err error) {
	if err != nil {
		JSON(w, statuscode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)

}

func JSON(w http.ResponseWriter, statuscode int, data interface{}) {
	w.WriteHeader(statuscode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
