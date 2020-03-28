package requestBody

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreate struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Thumbnail string `json:"thumbnail"`
}
