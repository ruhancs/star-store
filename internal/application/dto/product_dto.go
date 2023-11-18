package dto

import "star_store/internal/domain/entity"

type InputCreateProductDto struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	ZipCode     string `json:"zip_code"`
	Seller      string `json:"seller"`
	ThumbnailHD string `json:"thumbnail"`
}

type OutputCreateProductDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	ZipCode     string `json:"zip_code"`
	Seller      string `json:"seller"`
	ThumbnailHD string `json:"thumbnail"`
}

type OutputListProductDto struct {
	Items []*entity.Product
}

type OutputGetProductDto struct {
	Item *entity.Product
}
