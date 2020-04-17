package response

import (
	"todone/pkg/domain/entity"
)

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type UsersResponse []*UserResponse

func ConvertToUserResponse(userData *entity.User) *UserResponse {
	return &UserResponse{
		Name:      userData.Name,
		Thumbnail: userData.Thumbnail,
	}
}
