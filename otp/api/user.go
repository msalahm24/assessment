package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/msalahm24/otp/db/sqlc"
)

type createUserRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type createUserRes struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func (s *server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err := s.store.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			user, err := s.store.CreateUser(ctx, db.CreateUserParams{
				Name:        req.Name,
				PhoneNumber: req.PhoneNumber,
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
			}
			ctx.JSON(http.StatusCreated, createUserRes{
				ID:          user.ID,
				Name:        user.Name,
				PhoneNumber: user.PhoneNumber,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusBadRequest, errResponse(fmt.Errorf("Phone number already exists")))
}
