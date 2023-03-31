USE e_commerce_api;

CREATE TABLE IF NOT EXISTS `carts` (
    `id` VARCHAR(250) NOT NULL UNIQUE,
    `user_id` VARCHAR(250) NULL UNIQUE,
    `total_price` FLOAT NOT NULL DEFAULT 0,
    `date_created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),

    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS `cart_item` (
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
    `cart_id` VARCHAR(250) NOT NULL,
    `product_id` VARCHAR(250) NOT NULL DEFAULT '',
    `quantity` INT NOT NULL DEFAULT 1,
    `date_created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),

	
    FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY(cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    CONSTRAINT uq_cart_item UNIQUE (cart_id, product_id),
    CONSTRAINT pk_cart_item PRIMARY KEY (cart_id, product_id)
);