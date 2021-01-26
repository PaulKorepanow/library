package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string
}

func NewBook(title string) *Book {
	return &Book{
		Title: title,
	}
}
