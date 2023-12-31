// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID                int32            `json:"id"`
	Name              string           `json:"name"`
	PhoneNumber       string           `json:"phone_number"`
	Otp               pgtype.Text      `json:"otp"`
	OtpExpirationTime pgtype.Timestamp `json:"otp_expiration_time"`
}
