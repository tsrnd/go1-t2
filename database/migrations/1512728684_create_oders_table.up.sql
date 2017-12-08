CREATE SEQUENCE orders_seq;

CREATE TABLE IF NOT EXISTS orders (
  id INT NOT NULL DEFAULT NEXTVAL ('orders_seq'),
  total_price DOUBLE PRECISION NULL DEFAULT NULL,
  status INT NULL DEFAULT NULL,
  address VARCHAR(255) NULL DEFAULT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMP(0) NULL DEFAULT NULL,
  updated_at TIMESTAMP(0) NULL DEFAULT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_orders_users2
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
;

CREATE INDEX fk_orders_users2_idx ON orders (user_id ASC);
