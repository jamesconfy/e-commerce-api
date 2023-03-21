CREATE TABLE IF NOT EXISTS `users`(
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	`user_id` VARCHAR(250) NOT NULL UNIQUE,
    `first_name` VARCHAR(250) NULL,
    `last_name` VARCHAR(250) NULL,
    `email` VARCHAR(250) NOT NULL UNIQUE,
    `phone_number` VARCHAR(250) NOT NULL UNIQUE,
    `password` VARCHAR(250) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `access_token` VARCHAR(250) NOT NULL DEFAULT '',
    `refresh_token` VARCHAR(250) NOT NULL DEFAULT '',
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `address` VARCHAR(250) NOT NULL DEFAULT '',
    
    PRIMARY KEY(user_id)
);


CREATE TRIGGER `users_delete_trigger`
AFTER DELETE ON `users`
FOR EACH ROW
INSERT INTO `deleted_users`
	(`_id`, `user_id`, `first_name`, `last_name`, `email`, `phone_number`, `password`, `date_created`, `access_token`, `refresh_token`, `date_updated`, `address`, `date_deleted`)
VALUES
	(OLD._id, OLD.user_id, OLD.first_name, OLD.last_name, OLD.email, OLD.phone_number, OLD.password, OLD.date_created, OLD.access_token, OLD.refresh_token, OLD.date_updated, OLD.address, current_timestamp());
