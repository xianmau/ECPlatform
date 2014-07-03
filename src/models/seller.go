package models

type Seller struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"level"`
}
