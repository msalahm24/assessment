package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/msalahm24/otp/util"
)

type Store interface {
	Querier
	GenerateOTPtx(ctx context.Context, arg GenerateOTPtxParams) (GenerateOTPtxResult, error)
}

type SqlStore struct {
	*Queries
	db *pgx.Conn
}

type GenerateOTPtxParams struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type GenerateOTPtxResult struct {
	OTP            string `json:"OTP"`
	ExpirationTime string `json:"expiration_time"`
}

func NewStore(db *pgx.Conn) Store {
	return &SqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		rberr := tx.Rollback(ctx)
		if rberr != nil {
			return fmt.Errorf("tx error: %v, rbError: %v", err, rberr)
		}
		return err
	}
	return tx.Commit(ctx)
}

func (store *SqlStore) GenerateOTPtx(ctx context.Context, arg GenerateOTPtxParams) (GenerateOTPtxResult, error) {
	var result GenerateOTPtxResult
	var err error

	err = store.execTx(ctx, func(q *Queries) error {
		fmt.Println("Verify that the phone number exists")
		_, err := q.GetUserByPhoneNumber(ctx, arg.PhoneNumber)
		if err != nil {
			return err
		}
		//Generate opt
		otp := util.GenerateOTP()
		otpExpiration := time.Now().UTC().Add(1 * time.Minute)
		fmt.Println("Update OTP")
		user, err := q.UpdateOTP(ctx, UpdateOTPParams{
			PhoneNumber: arg.PhoneNumber,
			Otp: pgtype.Text{
				Valid:  true,
				String: otp,
			},
			OtpExpirationTime: pgtype.Timestamp{
				Valid: true,
				Time:  otpExpiration,
			},
		})
		if err != nil {
			return err
		}
		result.OTP = user.Otp.String
		result.ExpirationTime = user.OtpExpirationTime.Time.UTC().String()
		return nil
	})
	return result, err
}
