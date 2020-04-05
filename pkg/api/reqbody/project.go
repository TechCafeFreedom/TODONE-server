package reqbody

type ProjectCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
