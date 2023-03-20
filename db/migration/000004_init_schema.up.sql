CREATE TABLE IF NOT EXISTS `ratings`(
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
    `rating_id` VARCHAR(250) NOT NULL DEFAULT '',
    `product_id` VARCHAR(250) NOT NULL DEFAULT '',
    `user_id` VARCHAR(250) NOT NULL DEFAULT '',
    `rating` FLOAT NOT NULL DEFAULT 0.000,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),

    FOREIGN KEY(product_id) REFERENCES products(product_id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    PRIMARY KEY (rating_id)
);