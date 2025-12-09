package handler

import (
	"net/http"

	"github.com/Nucleussss/hikayat-forum-gateway/internal/dto"
	"github.com/Nucleussss/hikayat-forum-gateway/internal/service"
	postpb "github.com/Nucleussss/hikayat-proto/gen/go/post/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostHandler struct {
	PostServiceClient *service.PostServiceClient
}

func NewPostHandler(ps *service.PostServiceClient) *PostHandler {
	return &PostHandler{PostServiceClient: ps}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userIDInterface, exist := c.Get("user_id")
	if !exist {
		c.JSON(500, gin.H{"error": "Internal error: user ID not found in context"})
		return
	}

	// Assert the type of the user ID (should be string if set by JWT middleware)
	authorID, ok := userIDInterface.(string)
	if !ok {
		// Handle case where context value is not the expected type
		c.JSON(500, gin.H{"error": "Internal error: invalid user ID type in context"})
		return
	}

	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": "Invalid request body"})
		return
	}

	grpcReq := &postpb.CreatePostRequest{
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: authorID,
		Category: req.Category,
	}

	createdPost, err := h.PostServiceClient.CreatePost(c, grpcReq)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, createdPost)
}

func (h *PostHandler) GetPost(c *gin.Context) {

	var req dto.GetPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": "Failed to bind JSON"})
		return
	}

	grpcReq := &postpb.GetPostRequest{
		Id: req.ID,
	}

	post, err := h.PostServiceClient.GetPost(c, grpcReq)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot get post from service"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) ListPost(c *gin.Context) {
	// userIDInterface, exist := c.Get("user_id")
	// if !exist {
	// 	c.JSON(401, gin.H{"error": "Unauthorized"})
	// }

	// userID, ok := userIDInterface.(string)
	// if !ok {
	// 	c.JSON(500, gin.H{"error": "Invalid user ID"})
	// 	return
	// }

	var req dto.ListPostsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	grpcReq := &postpb.ListPostsRequest{
		AuthorId: req.AuthorID,
		Category: req.Category,
		Page:     req.Page,
		Limit:    req.Limit,
	}

	listPost, err := h.PostServiceClient.ListPost(c, grpcReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't get posts"})
		return
	}

	c.JSON(http.StatusOK, listPost)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	// Check if the user is authenticated and has a valid
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": "Invalid user ID type"})
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	grpcPost := &postpb.Post{
		Id:        req.Post.ID,
		Title:     req.Post.Title,
		Content:   req.Post.Content,
		AuthorId:  req.Post.AuthorID,
		Category:  req.Post.Category,
		CreatedAt: timestamppb.New(req.Post.CreatedAt),
		UpdatedAt: timestamppb.New(req.Post.UpdatedAt),
		IsDeleted: req.Post.IsDeleted,
	}

	grpcReq := &postpb.UpdatePostRequest{
		Id:         req.ID,
		Post:       grpcPost,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: req.UpdateMask},
		AuthorId:   userID,
	}

	listPost, err := h.PostServiceClient.UpdatePost(c, grpcReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't get posts"})
		return
	}

	c.JSON(http.StatusOK, listPost)

}

func (h *PostHandler) DeletePost(c *gin.Context) {
	// Check if the user is authenticated and get the user
	userIDInterface, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error ": "User ID is required"})
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": "Invalid user ID type"})
		return
	}

	// Get post ID from URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	grpcReq := &postpb.DeletePostRequest{
		Id:       id,
		AuthorId: userID,
	}

	_, err := h.PostServiceClient.DeletePost(c, grpcReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't get posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}
