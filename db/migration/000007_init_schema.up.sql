USE e_commerce_api;

CREATE TABLE IF NOT EXISTS auth (
    id VARCHAR(250) NOT NULL,
	user_id VARCHAR(250) NOT NULL,
    access_token VARCHAR(250) NOT NULL,
    refresh_token VARCHAR(250) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    
    PRIMARY KEY(user_id, access_token),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
