USE e_commerce_api;

CREATE TABLE IF NOT EXISTS `users`(
	`id` VARCHAR(250) NOT NULL UNIQUE,
    -- `cart_id` VARCHAR(250) NOT NULL UNIQUE DEFAULT '',
    `first_name` VARCHAR(250) NULL,
    `last_name` VARCHAR(250) NULL,
    `email` VARCHAR(250) NOT NULL UNIQUE,
    `phone_number` VARCHAR(250) NOT NULL UNIQUE,
    `password` VARCHAR(250) NOT NULL,
    `access_token` VARCHAR(250) NOT NULL DEFAULT '',
    `refresh_token` VARCHAR(250) NOT NULL DEFAULT '',
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `address` VARCHAR(250) NOT NULL DEFAULT '',
    
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS `deleted_users`(
	`id` VARCHAR(250),
    -- `cart_id` VARCHAR(250),
    `first_name` VARCHAR(250),
    `last_name` VARCHAR(250),
    `email` VARCHAR(250),
    `phone_number` VARCHAR(250),
    `password` VARCHAR(250),
    `access_token` VARCHAR(250),
    `refresh_token` VARCHAR(250),
    `date_created` TIMESTAMP,
    `date_updated` TIMESTAMP,
    `date_deleted` TIMESTAMP,
    `address` VARCHAR(250)
);


CREATE TRIGGER `users_delete_trigger`
AFTER DELETE ON `users`
FOR EACH ROW
INSERT INTO `deleted_users`
	(`id`, `first_name`, `last_name`, `email`, `phone_number`, `password`, `date_created`, `access_token`, `refresh_token`, `date_updated`, `address`, `date_deleted`)
VALUES
	(OLD.id, OLD.first_name, OLD.last_name, OLD.email, OLD.phone_number, OLD.password, OLD.date_created, OLD.access_token, OLD.refresh_token, OLD.date_updated, OLD.address, current_timestamp());
