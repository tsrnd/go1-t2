package models

import (
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
	id, quantityOld, priceOld := checkProductExitsOnCart(productId)
	if id > 0 {
		return Update(id, quantity+quantityOld, price+priceOld)
	}
	var cartDetail CartDetail
	cartDetail = CartDetail{CreatedAt: time.Now(), Price: price, Quantity: quantity, UserId: userId, ProductId: productId, OrderId: orderId}
	db.Create(&cartDetail)
	return cartDetail.Id, nil
}

// show cart
func ShowCart(idOrder interface{}) ([]*Order, error) {
	rows, err := db.Table("orders").Select("cart_details.id, cart_details.price as total_price, cart_details.quantity, products.name, products.image, products.price").Joins("JOIN cart_details ON orders.id=cart_details.order_id").Joins("JOIN products ON cart_details.product_id=products.id").Where("orders.id = ? AND orders.status = ?", idOrder, 0).Order("cart_details.id DESC").Rows()
	orders := make([]*Order, 0)
	for rows.Next() {
		order := new(Order)
		err := rows.Scan(&order.CartDetailId, &order.TotalPrice, &order.Quantity, &order.Name, &order.Image, &order.Price)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

// remove cart
func Remove(cartDetailId int64) (int64, error) {
	var cartDetail CartDetail
	cartDetail.Id = cartDetailId
	result := db.Delete(&cartDetail)
	if result.Error != nil {
		return 0, result.Error
	}
	return cartDetailId, nil
}

//update cart
func Update(cartDetailId int64, quantity int, price float64) (int64, error) {
	var cartDetail CartDetail
	cartDetail.Id = cartDetailId
	db.First(&cartDetail)
	cartDetail.Quantity = quantity
	cartDetail.Price = price
	result := db.Save(&cartDetail)
	return cartDetailId, result.Error
}

// check product exits cart
func checkProductExitsOnCart(idProduct int64) (int64, int, float64) {
	var cartDetail CartDetail
	db.Where(&CartDetail{ProductId: idProduct}).First(&cartDetail)
	return cartDetail.Id, cartDetail.Quantity, cartDetail.Price
}
