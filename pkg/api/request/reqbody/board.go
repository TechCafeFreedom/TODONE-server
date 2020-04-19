package reqbody

type BoardCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
