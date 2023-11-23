package factory

import (
	"star_store/internal/application/usecase"
	"star_store/internal/infra/repository"

	"github.com/gocql/gocql"
)

func CreateProductUseCaseFactory(db *gocql.Session) *usecase.CreateProducteUseCase {
	productRepository := repository.NewProductRepository(db)
	usecase := usecase.NewCreateProductUseCase(productRepository)
	return usecase
}

func ListProductsUseCaseFactory(db *gocql.Session) *usecase.ListProductUseCase {
	productRepository := repository.NewProductRepository(db)
	usecase := usecase.NewListProductUseCase(productRepository)
	return usecase
}

func GetProductUseCaseFactory(db *gocql.Session) *usecase.GetProductUseCase {
	productRepository := repository.NewProductRepository(db)
	usecase := usecase.NewGetProductUseCase(productRepository)
	return usecase
}

func CreateClientUseCaseFactory(db *gocql.Session) *usecase.CreateClietUseCase {
	clientRepository := repository.NewClientRepository(db)
	usecase := usecase.NewCreateClientUseCase(clientRepository)
	return usecase
}

func GetClientUseCaseFactory(db *gocql.Session) *usecase.GetClientUseCase {
	clientRepository := repository.NewClientRepository(db)
	usecase := usecase.NewGetClientUseCase(clientRepository)
	return usecase
}

func GetCartByClientIDUseCaseFactory(db *gocql.Session) *usecase.GetCartByClientIDUseCase {
	clientRepository := repository.NewCartRepository(db)
	usecase := usecase.NewGetCartByClientIDUseCase(clientRepository)
	return usecase
}

func InsertItemOnCartUseCaseFactory(db *gocql.Session) *usecase.InsertItemOnCartUseCase {
	productRepository := repository.NewProductRepository(db)
	cartRepository := repository.NewCartRepository(db)
	cartItemRepository := repository.NewCartItemRepository(db)
	usecase := usecase.NewInsertItemOnCartUseCase(cartRepository,cartItemRepository,productRepository)
	return usecase
}

func CheckoutUseCaseFactory(db *gocql.Session) *usecase.CheckoutUseCase {
	clientRepository := repository.NewClientRepository(db)
	cartRepository := repository.NewCartRepository(db)
	usecase := usecase.NewCheckoutUseCase(cartRepository,clientRepository)
	return usecase
}

func BuyUseCaseFactory(db *gocql.Session) *usecase.BuyUseCase {
	clientRepository := repository.NewClientRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	usecase := usecase.NewBuyUseCase(transactionRepository,clientRepository)
	return usecase
}
