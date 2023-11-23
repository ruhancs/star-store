package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"star_store/internal/application/dto"
	"star_store/internal/domain/entity"

	"github.com/go-chi/chi/v5"
)

//func (app *Application) CartExists(r *http.Request) bool {
//	return app.Session.Exists(r.Context(), "cart")
//}

func (app *Application) AddToCart(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.InputInserItemOnCartDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	product, err := app.GetProductUseCase.Execute(inputDto.ProductID)
	if err != nil {
		app.errorJson(w, err, http.StatusNotFound)
	}

	clientCart, err := app.GetCartByClientIDUseCase.Execute(inputDto.ClientID)
	if err != nil {
		fmt.Println("ERR TO CHECKCART IN SESSION")
		cart := entity.NewCart(inputDto.ClientID, []*entity.CartItem{})

		cartItem, err := entity.NewCartItem(product.Item.Title, inputDto.ClientID, cart.ID, inputDto.Quantity, product.Item.Price)
		cart.InsertItem(cartItem)

		output, err := app.InsertItemOnCartUseCase.Execute(inputDto, cart, cartItem)
		if err != nil {
			fmt.Println("ERR IN USECASE")
			app.errorJson(w, err, http.StatusInternalServerError)
		}

		app.writeJson(w, http.StatusOK, output)
		return
	}

	cart := &entity.Cart{
		ID:        clientCart.CartID,
		ClientID:  clientCart.ClientID,
		CartItems: clientCart.CartItems,
		Total:     clientCart.Total,
	}
	cartItem, _ := entity.NewCartItem(product.Item.Title, inputDto.ClientID, cart.ID, inputDto.Quantity, product.Item.Price)
	cart.InsertItem(cartItem)

	output, err := app.InsertItemOnCartUseCase.Execute(inputDto, cart, cartItem)
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
	}

	app.writeJson(w, http.StatusOK, output)
}

func (app *Application) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")
	cart, err := app.GetCartByClientIDUseCase.Execute(clientID)
	if err != nil {
		app.errorJson(w, errors.New("empty cart"))
		return
	}

	cartEntity := entity.NewCart(cart.ClientID, cart.CartItems)

	output, err := app.CheckoutUseCase.Execute(cartEntity)
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	app.writeJson(w, http.StatusOK, output)
}

func (app *Application) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.InputBuyUseCaseDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.PaymentErrMetric.Inc()
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	output, err := app.BuyUseCase.Execute(inputDto)
	if err != nil {
		app.PaymentErrMetric.Inc()
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	app.writeJson(w, http.StatusOK, output)
}
