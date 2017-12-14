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
	Db.QueryRow("INSERT INTO orders(created_at) VALUES($1) returning id;", time.Now()).Scan(&order.Id)
	return order.Id, nil
}

// func ShowOrder(idOrder int64) ([]Result, error) {
// 	var result []Result
// 	Db.Raw("SELECT cart_details.id, cart_details.price as total_price, cart_details.quantity, products.name, products.image, products.price FROM orders INNER JOIN cart_details ON orders.id=cart_details.order_id INNER JOIN products ON cart_details.product_id=products.id WHERE orders.id=? ORDER BY cart_details.id DESC", idOrder).Scan(&result)
// 	return result, nil
// }
func SetCurrentOrder(orderId interface{}, userId int64, totalPrice float64, address string) (int64, error) {
	update, err := Db.Prepare("update orders set total_price = $1, status = $2, user_id = $3 where id = $4")
	checkErr(err)
	res, err := update.Exec(totalPrice, 1, userId, orderId)
	checkErr(err)
	affect, err := res.RowsAffected()
	return affect, err
}
