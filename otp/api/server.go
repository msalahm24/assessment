package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/msalahm24/otp/db/sqlc"
)

type server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer will create a new http server and setup router
func NewServer(store db.Store) *server {
	server := &server{
		store: store,
	}

	server.setUpRouter()
	return server
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *server) Start(address string) error {
	return server.router.Run(address)
}

func (server *server) setUpRouter() {
	router := gin.Default()

	router.POST("/api/users", server.createUser)
	router.POST("/api/users/generateotp", server.generateOTP)
	router.POST("/api/users/verifyotp", server.verifyOTP)
	server.router = router
}
