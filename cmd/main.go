package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"star_store/internal/application/factory"
	"star_store/internal/domain/entity"
	"star_store/internal/infra/db"
	"star_store/internal/infra/web"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/client_golang/prometheus"
)

var paymentErrMetric = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "payment_errors",
	Help: "Total errors in payment route",
})

func init() {
	gob.Register(entity.Cart{}) //para armazer na usuario na sessao
}

func main() {
	cassandraConn, _ := db.ConnetToCassandraCluster()
	defer cassandraConn.Close()

	createProductUseCase := factory.CreateProductUseCaseFactory(cassandraConn)
	listProductUseCase := factory.ListProductsUseCaseFactory(cassandraConn)
	getProductUseCase := factory.GetProductUseCaseFactory(cassandraConn)
	createClientUseCase := factory.CreateClientUseCaseFactory(cassandraConn)
	getClientUseCase := factory.GetClientUseCaseFactory(cassandraConn)
	getCartByClientIDUseCase := factory.GetCartByClientIDUseCaseFactory(cassandraConn)
	insertItemOnCartUseCase := factory.InsertItemOnCartUseCaseFactory(cassandraConn)
	checkoutUseCase := factory.CheckoutUseCaseFactory(cassandraConn)
	buyUseCase := factory.BuyUseCaseFactory(cassandraConn)

	//session := initSession()
	var srv = &http.Server{}

	app := web.NewApplication(
		paymentErrMetric,
		//session,
		buyUseCase,
		checkoutUseCase,
		createClientUseCase,
		createProductUseCase,
		getCartByClientIDUseCase,
		getClientUseCase,
		getProductUseCase,
		insertItemOnCartUseCase,
		listProductUseCase,
		srv,
	)

	errChan := make(chan error)
	go app.GracefullyShutdown(errChan)
	if err := <-errChan; err != nil {
		fmt.Println(err)
	}
}

func initSession() *scs.SessionManager {
	gob.Register(entity.Cart{}) //para armazer na usuario na sessao
	session := scs.New()
	//informacoes da sessao sao armazenada no redis
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour //sessao dura 24h
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10, //tempo maximo para conexao
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	return redisPool
}
