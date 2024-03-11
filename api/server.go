package api

import (
	"medods-exercise/db"
	"medods-exercise/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config utils.Config
	router *gin.Engine
	store  *mongo.Database
}

func NewServer(config utils.Config) *Server {
	db := db.Setup()
	server := &Server{
		config: config,
		store:  db,
	}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
