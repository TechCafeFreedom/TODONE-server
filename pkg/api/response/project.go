package response

import "time"

type ProjectResponse struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

type ProjectsResponse []*ProjectResponse
