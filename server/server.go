package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"test-kulina/health"
	"test-kulina/users"
	"test-kulina/config"
)

// Start start the server
func Start() error {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	config := config.Get()

	userService := users.New()

	health.Register(router)
	users.Register(router, userService)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}

	go func() {
		log.Printf("Server starting on port :%d\n", config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return handleShutdown(srv)
}

func handleShutdown(srv *http.Server) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return err
}
