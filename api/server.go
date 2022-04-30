package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/imonasterio/go-sql-docker-aws/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs athe http server on a specific ip addresss
func (server *Server) Start(ipAddress string) error {
	return server.router.Run(ipAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error:": err.Error()}
}
