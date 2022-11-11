CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(100) NOT NULL);

INSERT INTO products(title, description)
VALUES ('Apple', 'Very tasty apple!'), ('Orange', 'Very orange and tasty orange!');
