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
	router.PATCH("/user/:id", server.updateUser)
	router.DELETE("/user/:id", server.deleteUser)
	router.POST("/category", server.createCategory)
	router.GET("/category/:id", server.GetCategory)
	router.GET("/categories", server.listAllCategory)
	router.GET("/categories/:id", server.listAllCategoryByUserID)
	router.PUT("/category/:id", server.updateCategory)
	router.DELETE("/category/:id", server.deleteCategory)

	server.router = router
	return *server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
