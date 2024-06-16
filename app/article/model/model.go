package model

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}
