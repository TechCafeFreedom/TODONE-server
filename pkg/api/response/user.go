package response

import "time"

type UserResponse struct {
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersResponse []*UserResponse
