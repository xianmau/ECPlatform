package models

type LinkCategory struct {
	Title    string `json:"title"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}
