package models

import (
	"fmt"
	"log"
	"time"
)

type CartDetail struct {
	Id        int64 `gorm:"primary_key"`
	UserId    int64
	Price     float64
	Quantity  int
	ProductId int64
	OrderId   int64
	CreatedAt time.Time
}
type Order struct {
	Id           int
	TotalPrice   float64
	Status       int
	Address      string
	CartDetailId int64
	Quantity     int
	Name         string
	Image        string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// insert cart detail
func InsertCartDetail(price float64, quantity int, userId int64, productId int64, orderId int64) (int64, error) {
	id, quantityOld, priceOld := checkProductExitsOnCart(productId, orderId)
	quantityF := float64(quantity)
	priceNew := price*quantityF + priceOld
	if id > 0 {
		return Update(id, quantity+quantityOld, priceNew)
	}
	var cartDetailId int64
	db.QueryRow("INSERT INTO cart_details(price, quantity, product_id, order_id, created_at) VALUES($1, $2, $3, $4, $5) returning id;", price, quantity, productId, orderId, time.Now()).Scan(&cartDetailId)
	return cartDetailId, nil
}

// show cart
func ShowCart(idOrder interface{}) ([]*Order, error) {
	carts := make([]*Order, 0)
	rows, _ := db.Query("select cart_details.id, cart_details.price as total_price, cart_details.quantity, products.name, products.image, products.price FROM orders INNER JOIN cart_details ON orders.id=cart_details.order_id INNER JOIN products ON cart_details.product_id=products.id where orders.id = $1 AND orders.status = $2 ORDER BY cart_details.id DESC", idOrder, 0)
	for rows.Next() {
		order := new(Order)
		err := rows.Scan(&order.CartDetailId, &order.TotalPrice, &order.Quantity, &order.Name, &order.Image, &order.Price)
		if err != nil {
			log.Fatal(err)
		}
		carts = append(carts, order)
	}
	return carts, nil
}

// remove cart
func Remove(cartDetailId int64) (int64, error) {
	delete, err := db.Prepare("delete from cart_details where id=$1")
	checkErr(err)
	res, err := delete.Exec(cartDetailId)
	checkErr(err)
	affect, err := res.RowsAffected()
	fmt.Println("id remove", affect)
	return affect, nil
}

//update cart
func Update(cartDetailId int64, quantity int, price float64) (int64, error) {
	update, err := db.Prepare("update cart_details set quantity = $1, price = $2 where id = $3")
	checkErr(err)
	res, err := update.Exec(quantity, price, cartDetailId)
	checkErr(err)
	affect, err := res.RowsAffected()
	fmt.Println("id update", affect)
	return affect, err
}

// check product exits cart
func checkProductExitsOnCart(idProduct int64, idOrder int64) (int64, int, float64) {
	var cartDetail CartDetail
	db.QueryRow("select cart_details.id, cart_details.quantity, cart_details.price FROM cart_details INNER JOIN orders ON orders.id = cart_details.order_id where cart_details.product_id = $1 AND orders.status = $2 AND cart_details.order_id = $3", idProduct, 0, idOrder).Scan(&cartDetail.Id, &cartDetail.Quantity, &cartDetail.Price)
	return cartDetail.Id, cartDetail.Quantity, cartDetail.Price
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
