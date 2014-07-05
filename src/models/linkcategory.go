package models

type LinkCategory struct {
	Name     string `json:"name"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}
