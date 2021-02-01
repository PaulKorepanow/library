package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
}

func NewBook(title string) *Book {
	return &Book{
		Title: title,
	}
}

func (b *Book) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b *Book) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, b); err != nil {
		return err
	}
	return nil
}
