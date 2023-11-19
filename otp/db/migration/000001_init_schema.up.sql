CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(50) UNIQUE NOT NULL,
    otp VARCHAR(10),
    otp_expiration_time TIMESTAMP
);