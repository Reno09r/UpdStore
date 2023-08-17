CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_fname VARCHAR(100) NOT NULL,
    user_lname VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    role_id INTEGER NOT NULL,
    hashed_password VARCHAR(100) NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

CREATE TABLE catalogs
(
    catalog_id SERIAL PRIMARY KEY,
    catalog_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE manufacturers
(
    manufacturer_id SERIAL PRIMARY KEY,
    manufacturer_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE products
(
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    manufacturer_id INTEGER,
    catalog_id INTEGER,
    FOREIGN KEY (catalog_id) REFERENCES catalogs (catalog_id),
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers (manufacturer_id)
);

CREATE TABLE price_change
(
    product_id INTEGER NOT NULL,
    date_price_change TIMESTAMP NOT NULL,
    new_price NUMERIC(9,2) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

CREATE TABLE purchases
(
    purchase_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    purchase_date TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE purchase_items
(
    purchase_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    product_count INTEGER NOT NULL,
    product_price NUMERIC(9,2) NOT NULL,
    CONSTRAINT PK_PURCHASE_ITEMS PRIMARY KEY (purchase_id, product_id),
    FOREIGN KEY (product_id) REFERENCES products (product_id),
    FOREIGN KEY (purchase_id) REFERENCES purchases (purchase_id)
);