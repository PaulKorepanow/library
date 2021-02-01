package nosqlbase

import (
	"github.com/stretchr/testify/assert"
	"main/internal/model"
	"testing"
)

func TestBookRepository_Create(t *testing.T) {
	testBase := TestDataBase(t)

	testBook := model.NewBook("New")
	err := testBase.BookRepository().Create(testBook)
	assert.NoError(t, err)
}

func TestBookRepository_FindByTitle(t *testing.T) {
	testBase := TestDataBase(t)

	testBook := model.NewBook("New")
	err := testBase.BookRepository().Create(testBook)
	assert.NoError(t, err)

	book, err := testBase.BookRepository().FindByTitle("New")
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "New", book.Title)
}
