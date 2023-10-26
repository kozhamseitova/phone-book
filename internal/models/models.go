package models

type Search struct {
	Name string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
}