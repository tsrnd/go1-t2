CREATE SEQUENCE products_seq;

CREATE TABLE IF NOT EXISTS products (
  id INT NOT NULL DEFAULT NEXTVAL ('products_seq'),
  name VARCHAR(45) NULL DEFAULT NULL,
  description TEXT NULL DEFAULT NULL,
  image VARCHAR(255) NULL DEFAULT NULL,
  price DOUBLE PRECISION NULL DEFAULT NULL,
  category_id INT NOT NULL,
  created_at TIMESTAMP(0) NULL DEFAULT NULL,
  updated_at TIMESTAMP(0) NULL DEFAULT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_products_categories
    FOREIGN KEY (category_id)
    REFERENCES categories (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
;

CREATE INDEX fk_products_categories_idx ON products (category_id ASC);