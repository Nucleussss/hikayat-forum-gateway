package gateway

import (
	"github.com/Nucleussss/hikayat-forum-gateway/internal/handler"
	"github.com/Nucleussss/hikayat-forum-gateway/internal/service"
	"github.com/Nucleussss/hikayat-forum-gateway/internal/transport"

	"log"
)

func main() {

	// Initialize gRPC clients
	userSvc, err := service.NewAuthServiceClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create user service client: %v", err)
	}

	postSvc, err := service.NewPostServiceClient("localhost:50052")
	if err != nil {
		log.Fatalf("Failed to create post service client: %v", err)
	}

	// Create handlers
	authHandler := handler.NewAuthHandler(userSvc)
	postHandler := handler.NewPostHandler(postSvc)

	// Setup and run HTTP server
	router := transport.SetupRouter(authHandler, postHandler)
	router.Run(":8080")
}
