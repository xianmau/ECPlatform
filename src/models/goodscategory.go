package models

type GoodsCategory struct {
	Name     string `json:"name"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}
