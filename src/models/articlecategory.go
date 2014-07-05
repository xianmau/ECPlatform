package models

type ArticleCategory struct {
	Name     string `json:"name"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}
