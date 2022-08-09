package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/sharvan/simplebank/db/sqlc"
	"github.com/sharvan/simplebank/token"
	"github.com/sharvan/simplebank/utils"
)

type Server struct {
	config     utils.Config
	Store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Coonot create Token")
	}
	server := &Server{Store: store, tokenMaker: tokenMaker, config: config}
	router := gin.Default()
	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)
	server.router = router
	return server, nil
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}

}
