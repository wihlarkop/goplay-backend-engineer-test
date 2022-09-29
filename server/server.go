package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Config struct {
	Address string
}

type Server struct {
	httpServer *http.Server
}

type GinHttpHandler struct {
	Server *Server
	Router *gin.Engine
}

func NewServer(handler http.Handler, config *Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Handler: handler,
			Addr:    fmt.Sprintf(":%s", config.Address),
		},
	}
}

func NewGinHttpRouter(config *Config) (*GinHttpHandler, error) {
	router := gin.New()

	router.Static("/media", "./media")
	router.Use(cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodDelete},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "X-App-Token", "X-Requested-With", "Authorization"},
	}))

	return &GinHttpHandler{
		Server: NewServer(router, config),
		Router: router,
	}, nil
}

func (g *GinHttpHandler) Start() {
	g.Server.StartServer()
}

func (g *Server) StartServer() {
	go func() {
		if err := g.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	log.Println("server is running ...")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err.Error())
		os.Exit(1)
	}

	log.Println("Server stopped.")
}
