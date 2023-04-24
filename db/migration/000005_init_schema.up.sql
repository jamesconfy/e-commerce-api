CREATE TABLE IF NOT EXISTS carts (
    id uuid DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
    user_id uuid NULL UNIQUE,
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS cart_item (
    _id SERIAL,
    cart_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	
    FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY(cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    CONSTRAINT uq_cart_item UNIQUE (cart_id, product_id),
    CONSTRAINT pk_cart_item PRIMARY KEY (cart_id, product_id)
);