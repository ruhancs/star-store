package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"star_store/internal/application/dto"

	"github.com/go-chi/chi/v5"
)

func(app *Application) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.InputCreateClientDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	output,err := app.CreateClientUseCase.Execute(inputDto)
	app.writeJson(w,http.StatusCreated,output)
}

func(app *Application) GetClientHandler(w http.ResponseWriter, r *http.Request) {
	clientId := chi.URLParam(r,"id")
	output,err := app.GetClientUseCase.Execute(clientId)
	if err != nil {
		app.errorJson(w,fmt.Errorf("client not found"),http.StatusNotFound)
		return
	}
	app.writeJson(w,http.StatusOK,output)
}