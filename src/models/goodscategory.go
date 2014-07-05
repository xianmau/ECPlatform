package models

type GoodsCategory struct {
	Title    string `json:"title"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}
