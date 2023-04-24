CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id uuid DEFAULT uuid_generate_v4(),
    first_name VARCHAR(250),
    last_name VARCHAR(250),
    email VARCHAR(250) NOT NULL UNIQUE,
    phone_number VARCHAR(250) NOT NULL UNIQUE,
    password VARCHAR(250) NOT NULL,
    address VARCHAR(250),
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS deleted_users (
	id uuid,
    first_name VARCHAR(250),
    last_name VARCHAR(250),
    email VARCHAR(250),
    phone_number VARCHAR(250),
    password VARCHAR(250),
    address VARCHAR(250),
    date_created TIMESTAMP,
    date_updated TIMESTAMP,
    date_deleted TIMESTAMP
);

CREATE OR REPLACE FUNCTION users_delete_function()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO deleted_users (id, first_name, last_name, email, phone_number, password, address, date_created, date_updated, date_deleted)
  VALUES (OLD.id, OLD.first_name, OLD.last_name, OLD.email, OLD.phone_number, OLD.password, OLD.address, OLD.date_created, OLD.date_updated, current_timestamp);
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users_delete_trigger
AFTER DELETE ON users
FOR EACH ROW
EXECUTE FUNCTION users_delete_function();
