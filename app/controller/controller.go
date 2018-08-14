package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Controller is a base type for all controllers
type Controller struct {
}

// SendJSON is used for returning a Json output
func (c *Controller) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Printf("Error while encoding JSON %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal Server Error"}`)
	}

	if v == nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"result": "data not found"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}
