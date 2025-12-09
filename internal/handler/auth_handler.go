package handler

import (
	"log"
	"net/http"

	"github.com/Nucleussss/hikayat-forum-gateway/internal/dto"
	"github.com/Nucleussss/hikayat-forum-gateway/internal/service"
	authpb "github.com/Nucleussss/hikayat-proto/gen/go/auth/v1"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthServiceClient *service.AuthServiceClient
}

func NewAuthHandler(as *service.AuthServiceClient) *AuthHandler {
	return &AuthHandler{AuthServiceClient: as}
}

// POST /login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	grpcReq := &authpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.AuthServiceClient.Login(c, grpcReq)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{
		"token":   res.Token,
		"message": res.Message,
	})
}

// POST /Register
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	grpcReq := &authpb.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.AuthServiceClient.Register(c, grpcReq)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{
		"message": res.Message,
	})
}

func (h *AuthHandler) GetUser(c *gin.Context) {
	var req dto.GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	grpcReq := &authpb.GetUserRequest{
		Id: req.Id,
	}

	user, err := h.AuthServiceClient.GetUser(c, grpcReq)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal error: failed to get user profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) UpdateUserProfile(c *gin.Context) {
	authenticatedUserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(500, gin.H{"error": "Internal error: user ID not found in context"})
	}

	var req dto.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	grpcReq := &authpb.UpdateUserProfileRequest{
		Id:   authenticatedUserID.(string),
		Name: req.Name,
	}

	res, err := h.AuthServiceClient.UpdateUserProfile(c, grpcReq)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{
		"message": res.Message,
	})

}

func (h *AuthHandler) ChangeUserEmail(c *gin.Context) {
	authenticatedUserID, exists := c.Get("user_id") // Or whatever key you used in middleware
	if !exists {
		// This shouldn't happen if AuthMiddleware is set up correctly for this route
		c.JSON(500, gin.H{"error": "Internal error: user ID not found in context"})
		return
	}

	var req dto.ChangeUserEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	grpcReq := &authpb.ChangeUserEmailRequest{
		Email: req.Email,
		Id:    authenticatedUserID.(string),
	}

	res, err := h.AuthServiceClient.ChangeUserEmail(c, grpcReq)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{
		"message": res.Message,
	})
}

func (h *AuthHandler) ChangeUserPassword(c *gin.Context) {
	authenticatedUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "Internal error: user ID not found in context"})
		return
	}

	var req dto.ChangeUserPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	grpcReq := &authpb.ChangeUserPasswordRequest{
		Id:              authenticatedUserID.(string),
		Currentpassword: req.CurrentPassword,
		Newpassword:     req.NewPassword,
	}

	res, err := h.AuthServiceClient.ChangeUserPassword(c, grpcReq)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{
		"message": res.Message,
	})
}

func (h *AuthHandler) DeleteUser(c *gin.Context) {
	authID, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "Internal error: user ID not found in context"})
		return
	}

	targetUserID, ok := authID.(string)
	if !ok {
		c.JSON(403, gin.H{"error": "Forbidden: You can only delete your own account."})
		return
	}

	// 4. Prepare the gRPC request using the target user ID from the path (or context if self-delete)
	grpcReq := &authpb.DeleteUserRequest{ // Use the correct proto package name (e.g., userpb)
		Id: targetUserID, // Use the ID you want to delete (here, the one from the path, verified above)
	}

	res, err := h.AuthServiceClient.DeleteUser(c, grpcReq) // Use the correct field name (e.g., h.userServiceClient)
	if err != nil {
		// TODO: Implement proper gRPC error code mapping to HTTP status codes
		// e.g., NotFound -> 404, Internal -> 500, etc.
		// The error message "Invalid email or password" is incorrect here.
		log.Printf("Error calling DeleteUser gRPC service: %v", err) // Log the actual error
		c.JSON(500, gin.H{"error": "Failed to delete user"})         // Generic message for client
		return
	}

	// 6. Respond
	c.JSON(200, gin.H{
		"message": res.Message, // Fixed typo: "massage" -> "message"
	})
}
