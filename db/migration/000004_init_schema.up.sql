USE e_commerce_api;

CREATE TABLE IF NOT EXISTS `ratings`(
    `product_id` VARCHAR(250) NOT NULL DEFAULT '',
    `user_id` VARCHAR(250) NOT NULL DEFAULT '',
    `rating` FLOAT NOT NULL DEFAULT 0.000,
    `date_created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `date_updated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),

    FOREIGN KEY(product_id) REFERENCES products(`id`) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(`id`) ON DELETE CASCADE,
    CONSTRAINT uq_rating_id UNIQUE (product_id, user_id),
    CONSTRAINT check_rating CHECK (rating <= 5.0 AND rating >= 0.0),
    CONSTRAINT pk_rating_id PRIMARY KEY (product_id, user_id)
);