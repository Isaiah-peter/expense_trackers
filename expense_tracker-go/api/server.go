package api

import (
	db "github.com/Isaiah-peter/expense_tracker/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/createuser", server.createUser)
	router.GET("/user/:id", server.GetUser)
	router.GET("/users", server.listAllUsers)

	server.router = router
	return *server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
