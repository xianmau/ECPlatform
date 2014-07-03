package models

type Role struct {
	Name      string   `json:"name"`
	Authority []string `json:"authority"`
}
