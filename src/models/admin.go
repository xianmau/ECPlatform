package models

type Admin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
