package web

import (
	"encoding/json"
	"net/http"
	"star_store/internal/application/dto"

	"github.com/go-chi/chi/v5"
)

func (app *Application) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.InputCreateProductDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	output,err := app.CreateProductUseCase.Execute(inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	app.writeJson(w,http.StatusCreated,output)
}

func(app *Application) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	output,err := app.ListProductsUseCase.Execute()
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
		return
	}
	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r,"id")
	output,err := app.GetProductUseCase.Execute(productID)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
		return
	}
	app.writeJson(w,http.StatusOK,output)
}

