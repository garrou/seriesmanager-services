package models

type Actor struct {
	Person struct {
		Id     int         `json:"id"`
		Name   string      `json:"name"`
		Shows  interface{} `json:"shows"`
		Movies interface{} `json:"movies"`
	} `json:"person"`
}
