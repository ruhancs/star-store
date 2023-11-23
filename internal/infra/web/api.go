package web

import (
	"log"
	"net/http"
	"star_store/internal/application/usecase"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	//"github.com/alexedwards/scs/v2"
)

type Application struct {
	PaymentErrMetric prometheus.Counter
	//Session                  *scs.SessionManager
	BuyUseCase      *usecase.BuyUseCase
	CheckoutUseCase *usecase.CheckoutUseCase
	//CreateCartUseCase        *usecase.CreateCartUseCase//
	CreateClientUseCase      *usecase.CreateClietUseCase
	CreateProductUseCase     *usecase.CreateProducteUseCase
	GetCartByClientIDUseCase *usecase.GetCartByClientIDUseCase //
	GetClientUseCase         *usecase.GetClientUseCase
	GetProductUseCase        *usecase.GetProductUseCase
	InsertItemOnCartUseCase  *usecase.InsertItemOnCartUseCase
	ListProductsUseCase      *usecase.ListProductUseCase
	SRV                      *http.Server
}

func NewApplication(
	paymentErrMetric prometheus.Counter,
	//session *scs.SessionManager,
	buyUseCase *usecase.BuyUseCase,
	checkoutUseCase *usecase.CheckoutUseCase,
	//createCartUseCase *usecase.CreateCartUseCase,
	createClientUseCase *usecase.CreateClietUseCase,
	createProductUseCase *usecase.CreateProducteUseCase,
	getCartByClientIDUseCase *usecase.GetCartByClientIDUseCase,
	getClientUseCase *usecase.GetClientUseCase,
	getProductUseCase *usecase.GetProductUseCase,
	insertItemOnCartUseCase *usecase.InsertItemOnCartUseCase,
	listProductsUseCase *usecase.ListProductUseCase,
	srv *http.Server,
) *Application {
	return &Application{
		PaymentErrMetric: paymentErrMetric,
		//Session:                  session,
		BuyUseCase:      buyUseCase,
		CheckoutUseCase: checkoutUseCase,
		//CreateCartUseCase:        createCartUseCase,
		CreateClientUseCase:      createClientUseCase,
		CreateProductUseCase:     createProductUseCase,
		GetCartByClientIDUseCase: getCartByClientIDUseCase,
		GetClientUseCase:         getClientUseCase,
		GetProductUseCase:        getProductUseCase,
		InsertItemOnCartUseCase:  insertItemOnCartUseCase,
		ListProductsUseCase:      listProductsUseCase,
		SRV:                      srv,
	}
}

func (app *Application) Server() error {
	app.SRV.Addr = ":8000"
	app.SRV.Handler = app.routes()
	app.SRV.IdleTimeout = 30 * time.Second
	app.SRV.ReadTimeout = 1 * time.Second
	app.SRV.ReadHeaderTimeout = 1 * time.Second
	app.SRV.WriteTimeout = 5 * time.Second

	log.Println("Runing server on port 8000...")
	return app.SRV.ListenAndServe()
}
