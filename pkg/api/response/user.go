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

func ConvertToUsersResponse(userSlice entity.UserSlice) UsersResponse {
	res := make(UsersResponse, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserResponse(userData))
	}
	return res
}

func ConvertToUserResponse(userData *entity.User) *UserResponse {
	// nilチェック
	if userData == nil {
		return nil
	}

	return &UserResponse{
		ID:        userData.ID,
		Name:      userData.Name,
		Thumbnail: userData.Thumbnail,
	}
}
