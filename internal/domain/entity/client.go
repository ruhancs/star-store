package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	Number string `json:"card_number"`
	CVV int `json:"cvv"`
	CardHolderName string `json:"card_holder_name"`
}

type Client struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ZipCode string `json:"zip_code"`
}

func NewClient(name,zipcode string)(*Client,error){
	client := &Client{
		ID: uuid.NewV4().String(),
		Name: name,
		ZipCode: zipcode,
	}
	err := client.isValid()
	if err != nil {
		return nil,err
	}
	return client,nil
}

func(c *Client) isValid() error {
	_,err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}