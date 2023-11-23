package repository

import (
	"star_store/internal/domain/entity"

	"github.com/gocql/gocql"
)

type ClientRepository struct {
	DB *gocql.Session
}

func NewClientRepository(db *gocql.Session) *ClientRepository{
	return &ClientRepository{
		DB: db,
	}
}

func(c *ClientRepository) Create(client *entity.Client) error {
	err := c.DB.Query(`INSERT INTO clients(id, name, zip_code) VALUES (?, ?, ?);`,client.ID,client.Name,client.ZipCode).Exec()
	if err != nil {
		return err
	}

	return nil
}

func(c *ClientRepository) Get(id string) (*entity.Client,error) {
	var client entity.Client
	err := c.DB.Query("SELECT id,name,zip_code from clients WHERE id=?;",id).Scan(&client.ID,&client.Name,&client.ZipCode)
	if err != nil {
		return nil,err
	}

	return &client,nil
}

func(c *ClientRepository) Delete(id string) error {
	err := c.DB.Query("DELETE from clients WHERE id=?;",id).Exec()
	if err != nil {
		return err
	}

	return nil
}