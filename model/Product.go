package model

type Product struct {
	Id          int
	Title       string `json:"Title" db:"title" validate:"required,min=2"`
	Description string `json:"Description" db:"description" validate:"required,min=2"`
}
