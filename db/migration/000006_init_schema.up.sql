CREATE TYPE status_type AS ENUM('ACTIVE', 'ONGOING', 'COMPLETED', 'FAILED');
CREATE TYPE payment_method_type AS ENUM('CARD', 'PAY-ON-DELIVERY', 'PICKUP-STATION');

CREATE TABLE IF NOT EXISTS checkout (
	id uuid DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
    quantity INT NOT NULL,
    cart_id uuid NOT NULL,
    product_id uuid NOT NULL,
    status status_type DEFAULT 'ACTIVE', 
    payment_method payment_method_type DEFAULT 'CARD', 
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);