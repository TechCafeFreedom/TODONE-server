package reqheader

type ProjectLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProjectCreate struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Thumbnail string `json:"thumbnail"`
}

type ProjectGet struct {
	Authorization string
}
