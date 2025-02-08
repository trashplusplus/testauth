CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Products (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    price INTEGER NOT NULL
);

INSERT INTO Products (title, price) VALUES
    ('Citizen', 5000),
    ('Casio', 3000),
    ('Timex', 3000),
    ('Seiko', 8000),
    ('Panerai', 100000),
    ('AP', 2000000),
    ('Brew', 15000),
    ('IWC', 500000),
    ('Certina', 80000),
    ('Spinnaker', 10000)
ON CONFLICT DO NOTHING;
