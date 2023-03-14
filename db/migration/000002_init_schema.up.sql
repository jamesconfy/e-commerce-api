CREATE TABLE IF NOT EXISTS `users`(
    `_id` INT NOT NULL AUTO_INCREMENT UNIQUE,
	`user_id` VARCHAR(250) NOT NULL UNIQUE,
    `first_name` VARCHAR(250) NULL,
    `last_name` VARCHAR(250) NULL,
    `email` VARCHAR(250) NOT NULL UNIQUE,
    `phone_number` VARCHAR(250) NOT NULL UNIQUE,
    `password` VARCHAR(250) NOT NULL, 
    `date_created` VARCHAR(250) NOT NULL,
    `date_updated` VARCHAR(250) NOT NULL DEFAULT '',
    `address` VARCHAR(250) NOT NULL DEFAULT '',
    
    PRIMARY KEY(user_id)
);