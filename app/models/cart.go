package models

<<<<<<< Updated upstream
import "time"

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
	cartDetail = CartDetail{CreatedAt: time.Now(), Price: price, Quantity: quantity, UserId: userId, ProductId: productId, OrderId: orderId}
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
=======
import (
	"time"

	"github.com/jinzhu/gorm"
)

type Orders struct {
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

type Order struct {
	gorm.Model
	TotalPrice float64
	Status     int
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// insert order
func InsertOrder() int {
	var orderInsert = Orders{Address: "NGUYEN VAN A", UserId: 1}
	db.Create(&orderInsert)
	return orderInsert.Id
}

// insert cart detail
// func InsertCartDetail(price float64, quantity int, userId int, productId int, orderId interface{}) (int64, error) {
// 	var orderDetailInsert = Orders{Address: "NGUYEN VAN A", UserId: 1}
// 	stmtCartDetail, _ := db.Prepare("INSERT cart_details SET price=?,quantity=?,user_id=?,product_id=?,order_id=?,created_at=?")
// 	resCartDetail, _ := stmtCartDetail.Exec(price, quantity, userId, productId, orderId, time.Now())
// 	idCartDetail, _ := resCartDetail.LastInsertId()
// 	// defer stmtCartDetail.Close()
// 	return idCartDetail, nil
// }

// show cart
func ShowCart(idOrder interface{}) ([]*Orders, error) {
	rows, err := db.Table("cart_details").Select("cart_details.id, cart_details.price as total_price, cart_details.quantity, products.name, products.image, products.price").Joins("INNER JOIN cart_details ON orders.id=cart_details.order_id INNER JOIN products ON cart_details.product_id=products.id").Where("orders.id = ? AND orders.status = ?", idOrder, 0).Order("cart_details.id DESC").Rows()
	orders := make([]*Orders, 0)
	for rows.Next() {
		order := new(Orders)
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
// func Remove(orderId interface{}, cartDetailId int) (int64, error) {
// 	delete, err := db.Prepare("delete from cart_details where id=? AND order_id=?")
// 	if err != nil {
// 		return 0, err
// 	}
// 	result, err := delete.Exec(cartDetailId, orderId)
// 	if err != nil {
// 		return 0, err
// 	}
// 	affect, err := result.RowsAffected()
// 	return affect, err
// }

//update cart
// func Update(cartDetailId int, quantity int, totalPrice float64) (int64, error) {
// 	update, err := db.Prepare("update cart_details set quantity=?, price=? where id=?")
// 	if err != nil {
// 		return 0, err
// 	}
// 	res, err := update.Exec(quantity, totalPrice, cartDetailId)
// 	if err != nil {
// 		return 0, err
// 	}
// 	affect, err := res.RowsAffected()
// 	return affect, err
// }
>>>>>>> Stashed changes
