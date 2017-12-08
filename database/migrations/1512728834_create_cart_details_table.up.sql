CREATE SEQUENCE cart_details_seq;

CREATE TABLE IF NOT EXISTS cart_details (
  id INT NOT NULL DEFAULT NEXTVAL ('cart_details_seq'),
  price DOUBLE PRECISION NULL DEFAULT NULL,
  quantity INT NULL DEFAULT NULL,
  user_id INT NOT NULL,
  product_id INT NOT NULL,
  order_id INT NOT NULL,
  updated_at TIMESTAMP(0) NULL DEFAULT NULL,
  created_at TIMESTAMP(0) NULL DEFAULT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_orders_users1
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_orders_products1
    FOREIGN KEY (product_id)
    REFERENCES products (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT fk_carts_orders1
    FOREIGN KEY (order_id)
    REFERENCES orders (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
;

CREATE INDEX fk_orders_users1_idx ON cart_details (user_id ASC);
CREATE INDEX fk_orders_products1_idx ON cart_details (product_id ASC);
CREATE INDEX fk_carts_orders1_idx ON cart_details (order_id ASC);