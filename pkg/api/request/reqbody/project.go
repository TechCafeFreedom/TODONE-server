package reqbody

type ProjectCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectPK struct {
	ID int `json:"id"`
}
