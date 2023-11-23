package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func (app *Application) writeJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out,err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key,value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	_,err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) errorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	log.Println(err)
	
	return app.writeJson(w,statusCode,payload)
}