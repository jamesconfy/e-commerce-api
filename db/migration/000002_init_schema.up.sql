USE e_commerce_api;

CREATE TABLE IF NOT EXISTS `users`(
	`id` VARCHAR(250) NOT NULL UNIQUE,
    `first_name` VARCHAR(250) NULL,
    `last_name` VARCHAR(250) NULL,
    `email` VARCHAR(250) NOT NULL UNIQUE,
    `phone_number` VARCHAR(250) NOT NULL UNIQUE,
    `password` VARCHAR(250) NOT NULL,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS `deleted_users`(
	`id` VARCHAR(250),
    `first_name` VARCHAR(250),
    `last_name` VARCHAR(250),
    `email` VARCHAR(250),
    `phone_number` VARCHAR(250),
    `password` VARCHAR(250),
    `date_created` TIMESTAMP,
    `date_updated` TIMESTAMP,
    `date_deleted` TIMESTAMP
);


CREATE TRIGGER `users_delete_trigger`
AFTER DELETE ON `users`
FOR EACH ROW
INSERT INTO `deleted_users`
	(`id`, `first_name`, `last_name`, `email`, `phone_number`, `password`, `date_created`, `date_updated`, `date_deleted`)
VALUES
	(OLD.id, OLD.first_name, OLD.last_name, OLD.email, OLD.phone_number, OLD.password, OLD.date_created, OLD.date_updated, current_timestamp());
