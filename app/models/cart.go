package models

import (
	"time"
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

// insert order
func InsertOrder() (int64, error) {
	stmt, _ := db.Prepare("INSERT orders SET address=?,user_id=?,created_at=?")
	res, _ := stmt.Exec("tp đà nẵng", "1", time.Now())
	id, _ := res.LastInsertId()
	defer stmt.Close()
	return id, nil
}

// insert cart detail
func InsertCartDetail(price float64, quantity int, userId int, productId int, orderId interface{}) (int64, error) {
	stmtCartDetail, _ := db.Prepare("INSERT cart_details SET price=?,quantity=?,user_id=?,product_id=?,order_id=?,created_at=?")
	resCartDetail, _ := stmtCartDetail.Exec(price, quantity, userId, productId, orderId, time.Now())
	idCartDetail, _ := resCartDetail.LastInsertId()
	// defer stmtCartDetail.Close()
	return idCartDetail, nil
}

// show cart
func ShowCart(idOrder interface{}) ([]*Orders, error) {
	rows, err := db.Query("SELECT cart_details.id, cart_details.price as total_price, cart_details.quantity, products.name, products.image, products.price FROM orders INNER JOIN cart_details ON orders.id=cart_details.order_id INNER JOIN products ON cart_details.product_id=products.id WHERE orders.id=? AND orders.status=? ORDER BY cart_details.id DESC", idOrder, 0)
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
func Remove(orderId interface{}, cartDetailId int) (int64, error) {
	delete, err := db.Prepare("delete from cart_details where id=? AND order_id=?")
	if err != nil {
		return 0, err
	}
	result, err := delete.Exec(cartDetailId, orderId)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

//update cart
func Update(cartDetailId int, quantity int, totalPrice float64) (int64, error) {
	update, err := db.Prepare("update cart_details set quantity=?, price=? where id=?")
	if err != nil {
		return 0, err
	}
	res, err := update.Exec(quantity, totalPrice, cartDetailId)
	if err != nil {
		return 0, err
	}
	affect, err := res.RowsAffected()
	return affect, err
}
