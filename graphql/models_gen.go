// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type NewPost struct {
	Image   *graphql.Upload `json:"image"`
	Title   string          `json:"title"`
	Summary string          `json:"summary"`
	Link    string          `json:"link"`
}

type Post struct {
	ID       string    `json:"id"`
	ImageURL string    `json:"imageURL"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Link     string    `json:"link"`
	PostedOn time.Time `json:"postedOn"`
}
