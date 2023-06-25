package model

import "time"

type Person struct {
	Id       int       `json:"id" query:"id"`
	Name     string    `json:"name"`
	Phone    int       `json:"phone"`
	Birthday time.Time `json:"birthday"`
}
