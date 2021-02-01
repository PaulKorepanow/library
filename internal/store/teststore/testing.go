package teststore

import (
	"main/internal/model"
	"testing"
)

func NewTestBook(t *testing.T) *model.Book {
	return &model.Book{
		Title: "Harry Potter",
	}
}
