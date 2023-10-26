package entity

import (
	"errors"
	"log"
)

type Order struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID: id,
		Price: price,
		Tax: tax,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (order *Order) Validate() error {
	if order.ID == "" {
		return errors.New("ID is required")
	}
	if (order.Price <= 0) {
		return errors.New("Price must be positive")
	}
	return nil
}

func (order *Order) CalculateFinalPrice() (*Order, error) {
	err := order.Validate()
	if err != nil {
		return nil, errors.New("Impossible Calculate Final Price: Validation failed")
	}
	// Calculate final price using the percentage tax over the price
	order.FinalPrice = order.Price + (order.Price * (order.Tax / 100))
	
	log.Println("Final Price: ", order.FinalPrice)
	log.Println("Tax Order: ", order.Tax)
	log.Println("Price: ", order.Price)
	return order, nil
}