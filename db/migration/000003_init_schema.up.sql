CREATE TABLE IF NOT EXISTS `products`(
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
    `product_id` VARCHAR(250) NOT NULL,
    `user_id` VARCHAR(250) NOT NULL,
    `name` VARCHAR(250) NOT NULL,
    `description` VARCHAR(250) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `image` VARCHAR(250) NULL,

    FOREIGN KEY (`user_id`) REFERENCES e_commerce_api.users(`user_id`) ON DELETE CASCADE,
    PRIMARY KEY(`product_id`)
);