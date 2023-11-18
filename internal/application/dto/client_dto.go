package dto

import "star_store/internal/domain/entity"

type InputCreateClientDto struct {
	Name    string `json:"name"`
	ZipCode string `json:"zip_code"`
}

type OutputCreateClientDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ZipCode string `json:"zip_code"`
}

type OutputGetClientDto struct {
	Client *entity.Client
}
