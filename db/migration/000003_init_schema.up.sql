CREATE TABLE IF NOT EXISTS `products`(
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
    `product_id` VARCHAR(250) NOT NULL,
    `user_id` VARCHAR(250) NOT NULL,
    `name` VARCHAR(250) NOT NULL,
    `description` VARCHAR(250) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `image` VARCHAR(250) NULL,
    `price` FLOAT NOT NULL DEFAULT 0.00,

    FOREIGN KEY (`user_id`) REFERENCES e_commerce_api.users(`user_id`) ON DELETE CASCADE,
    PRIMARY KEY(`product_id`)
);

CREATE TRIGGER `products_delete_trigger`
AFTER DELETE ON `products`
FOR EACH ROW
INSERT INTO `deleted_products`
	(`_id`, `product_id`, `user_id`, `name`, `description`, `date_created`, `date_updated`, `image`, `price`, `date_deleted`)
VALUES
	(OLD._id, OLD.product_id, OLD.user_id, OLD.name, OLD.description, OLD.date_created, OLD.date_updated, OLD.image, OLD.price, current_timestamp());
