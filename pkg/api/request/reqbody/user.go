package reqbody

type UserCreate struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}
