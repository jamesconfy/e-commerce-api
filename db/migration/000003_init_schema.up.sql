CREATE TABLE IF NOT EXISTS products (
    id uuid DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
    user_id uuid NOT NULL,
    name VARCHAR(250) NOT NULL,
    description VARCHAR(250) NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    image VARCHAR(250) NULL,
    price FLOAT NOT NULL DEFAULT 0.00,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY(id),
    CONSTRAINT chk_price CHECK (price >= 0.00)
);

CREATE TABLE IF NOT EXISTS deleted_products (
    id uuid,
    user_id uuid,
    name VARCHAR(250),
    description VARCHAR(250),
    date_created TIMESTAMP,
    date_updated TIMESTAMP,
    date_deleted TIMESTAMP,
    image VARCHAR(250),
    price FLOAT
);

CREATE OR REPLACE FUNCTION product_delete_function()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO deleted_products(id, user_id, name, description, date_created, date_updated, image, price, date_deleted)
  VALUES(OLD.id, OLD.user_id, OLD.name, OLD.description, OLD.date_created, OLD.date_updated, OLD.image, OLD.price, current_timestamp);
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER product_delete_trigger
AFTER DELETE ON products
FOR EACH ROW
EXECUTE FUNCTION product_delete_function();