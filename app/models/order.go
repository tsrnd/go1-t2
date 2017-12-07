package models

import (
	"time"
)

type Orders struct {
	Id         int64 `gorm:"primary_key"`
	TotalPrice float64
	Status     int
	Address    string
	UserId     int64 `gorm:"default: null"`
	CreatedAt  time.Time
}
type Result struct {
	Id           int
	TotalPrice   float64
	Status       int
	Address      string
	UserId       int
	CartDetailId int
	Price        float64
	Quantity     int
	ProductId    int
	OrderId      int
	Name         string
	Image        string
}

// insert order
func InsertOrder(userId int64) (int64, error) {
	var order Orders
	order = Orders{Address: "Da Nang City", CreatedAt: time.Now()}
	if userId > 0 {
		order.UserId = userId
	}
	db.Create(&order)
	return order.Id, nil
}

func ShowOrder(idOrder int64) ([]Result, error) {
	var result []Result
	db.Raw("SELECT cart_details.id, cart_details.price as total_price, cart_details.quantity, products.name, products.image, products.price FROM orders INNER JOIN cart_details ON orders.id=cart_details.order_id INNER JOIN products ON cart_details.product_id=products.id WHERE orders.id=? ORDER BY cart_details.id DESC", idOrder).Scan(&result)
	return result, nil
}
func SetCurrentOrder(orderId, userId int64) {
	var order Orders
	order = Orders{Id: orderId}
	db.First(&order)
	order.UserId = userId
	db.Save(&order)
}
