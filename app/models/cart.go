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

// insert cart detail
func InsertCartDetail(price float64, quantity int, userId int64, productId int64, orderId int64) (int64, error) {
	var cartDetail CartDetail
	cartDetail = CartDetail{CreatedAt: time.Now(), Price: price, Quantity: quantity, ProductId: productId, OrderId: orderId}
	if userId > 0 {
		cartDetail.UserId = userId
	}
	db.Create(&cartDetail)
	return cartDetail.Id, nil
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
