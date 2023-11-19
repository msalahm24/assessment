-- name: CreateUser :one
INSERT INTO users (
    name,
    phone_number
) VALUES (
    $1,$2
) 
RETURNING *;

-- name: GetOTP :one
SELECT otp, otp_expiration_time
FROM users
WHERE phone_number = $1
LIMIT 1;

-- name: UpdateOTP :one
UPDATE users 
SET otp = $2, otp_expiration_time = $3 
WHERE phone_number = $1
RETURNING *;

-- name: GetUserByPhoneNumber :one
SELECT id, name, phone_number, otp, otp_expiration_time
FROM users
WHERE phone_number = $1;