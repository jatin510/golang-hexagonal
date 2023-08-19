package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jatin510/golang-hexagonal/internal/controller"
	"github.com/jatin510/golang-hexagonal/internal/core/server"
	"github.com/jatin510/golang-hexagonal/internal/core/service"
	"github.com/jatin510/golang-hexagonal/internal/infra/config"
	"github.com/jatin510/golang-hexagonal/internal/infra/repository"
)

func main() {
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())
  
	// Initialize the database connection
	db, err := repository.NewDB(
	 config.DatabaseConfig{
	  Driver: "mysql",
	  Url: "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=true&loc=UTC&tls=false&readTimeout=3s&writeTimeout=3s&timeout=3s&clientFoundRows=true",
	  ConnMaxLifetimeInMinute: 3,
	  MaxOpenConns:            10,
	  MaxIdleConns:            1,
	 },
	)
	if err != nil {
	 log.Fatalf("failed to new database err=%s\n", err.Error())
	}
  
	// Create the UserRepository
	userRepo := repository.NewUserRepository(db)
  
	// Create the UserService
	userService := service.NewUserService(userRepo)
	
	// Create the UserController
	userController := controller.NewUserController(instance, userService)
	
	// Initialize the routes for UserController
	userController.InitRouter()
  
	// Create the HTTP server
	httpServer := server.NewHttpServer(
	 instance,
	 config.HttpServerConfig{
	  Port: 8000,
	 },
	)
		
	// Start the HTTP server
	httpServer.Start()
	defer httpServer.Stop()
  
	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
	 c,
	 os.Interrupt,
	 syscall.SIGHUP,
	 syscall.SIGINT,
	 syscall.SIGQUIT,
	 syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
  }