package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/msalahm24/otp/db/sqlc"
)

type generateOTPReq struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type generateOTPRes struct {
	OTP            string `json:"OTP"`
	ExpirationTime string `json:"expiration_time"`
}

func (s *server) generateOTP(ctx *gin.Context) {
	var req generateOTPReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	result, err := s.store.GenerateOTPtx(ctx, db.GenerateOTPtxParams{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(fmt.Errorf("User Not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, generateOTPRes{
		ExpirationTime: result.ExpirationTime,
		OTP:            result.OTP,
	})
}

type verifyOTPReq struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	OTP         string `json:"OTP" binding:"required"`
}

type verifyOTPRes struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (s *server) verifyOTP(ctx *gin.Context) {
	var req verifyOTPReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	otp, err := s.store.GetOTP(ctx, req.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(fmt.Errorf("User Not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	if time.Now().UTC().Before(otp.OtpExpirationTime.Time) {
		if otp.Otp.String == req.OTP {
			ctx.JSON(http.StatusOK, verifyOTPRes{
				Status: "Verified",
			})
			return
		}
		ctx.JSON(http.StatusOK, verifyOTPRes{
			Status: "Not Verified",
			Error:  "Not correct OTP",
		})
		return
	}

	ctx.JSON(http.StatusOK, verifyOTPRes{
		Status: "Not Verified",
		Error:  "Expired OTP",
	})
	return
}
