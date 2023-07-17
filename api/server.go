package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/korvised/go-simplebank/db/sqlc"
	"github.com/korvised/go-simplebank/token"
	"github.com/korvised/go-simplebank/util"
)

// Server serves server HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer create a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can not crate token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil, fmt.Errorf("can not register currency validate: %w", err)
		}
	}

	// mount routes
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoute := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoute.POST("/accounts", server.createAccount)
	authRoute.GET("/accounts/:id", server.getAccount)
	authRoute.GET("/accounts", server.listAccount)

	authRoute.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
