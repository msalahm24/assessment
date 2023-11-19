// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetOTP(ctx context.Context, phoneNumber string) (GetOTPRow, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
	UpdateOTP(ctx context.Context, arg UpdateOTPParams) (User, error)
}

var _ Querier = (*Queries)(nil)
