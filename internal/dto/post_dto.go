package dto // Or your preferred package name

import (
	"time"
)

// --- Models ---

// Post represents the data structure for a blog post.
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"author_id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"` // Assuming time.Time for simplicity
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// --- Requests ---

// CreatePostRequest holds the data needed to create a new post.
// HTTP Method: POST
// URL: /api/v1/posts
type CreatePostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"author_id"` // Renamed to match Go convention
	Category string `json:"category"`
}

// GetPostRequest holds the data needed to retrieve a single post.
// HTTP Method: GET
// URL: /api/v1/posts/{id} (where {id} is the post ID)
type GetPostRequest struct {
	ID string `json:"id"`
}

// ListPostsRequest holds the data needed to list posts with optional filters and pagination.
// HTTP Method: GET
// URL: /api/v1/posts (with query parameters for filters and pagination)
// Example: /api/v1/posts?author_id=123&category=general&page=1&limit=10
type ListPostsRequest struct {
	AuthorID string `json:"author_id,omitempty"` // Optional filter
	Category string `json:"category,omitempty"`  // Optional filter
	Page     int32  `json:"page,omitempty"`      // Optional, defaults might be handled by logic
	Limit    int32  `json:"limit,omitempty"`     // Optional, defaults might be handled by logic
}

// UpdatePostRequest holds the data needed to update an existing post.
// HTTP Method: PATCH (or PUT)
// URL: /api/v1/posts/{id} (where {id} is the post ID)
// PATCH is often preferred for partial updates.
// PUT would replace the entire resource.
type UpdatePostRequest struct {
	ID         string   `json:"id"`
	Post       Post     `json:"post"`        // The updated post data
	UpdateMask []string `json:"update_mask"` // Placeholder: FieldMask is complex in Go, often handled differently or as a []string of field names
	// Example: UpdateMask []string `json:"update_mask,omitempty"`
}

// DeletePostRequest holds the data needed to delete a post.
// HTTP Method: DELETE
// URL: /api/v1/posts/{id} (where {id} is the post ID)
type DeletePostRequest struct {
	ID string `json:"id"`
}

// --- Responses ---

// ListPostsResponse holds the data returned when listing posts.
type ListPostsResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"has_more"`
}

// GetPostResponse (Added for completeness)
// This would be the response for a GET /api/v1/posts/{id} request.
type GetPostResponse struct {
	Post Post `json:"post"`
}

// CreatePostResponse (Added for completeness)
// This would be the response for a POST /api/v1/posts request.
type CreatePostResponse struct {
	Post Post `json:"post"`
}

// UpdatePostResponse (Added for completeness)
// This would be the response for a PATCH /api/v1/posts/{id} request.
type UpdatePostResponse struct {
	Post Post `json:"post"`
}

// DeletePostResponse (Added for completeness)
// This would be the response for a DELETE /api/v1/posts/{id} request.
// Often, DELETE responses are 204 No Content, but this struct represents a scenario where you might return metadata.
type DeletePostResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
