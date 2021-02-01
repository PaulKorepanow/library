package sqlbase

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"main/internal/model"
	"testing"
)

func TestBookRepository_Create(t *testing.T) {
	db, close := TestDatabase(t)
	defer close("books")

	testBook := model.NewBook("NewBook")
	err := db.BookRepository().Create(testBook)
	assert.NoError(t, err)
}

func TestBookRepository_FindBookByTitle(t *testing.T) {
	db, close := TestDatabase(t)
	defer close("books")

	testBook := model.NewBook("NewBook")
	err := db.BookRepository().Create(testBook)
	assert.NoError(t, err)

	book, err := db.BookRepository().FindBookByTitle("NewBook")
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, book.Title, "NewBook")
}

func TestBookRepository_DeleteBookByID(t *testing.T) {
	db, close := TestDatabase(t)
	defer close("books")

	testBook := model.NewBook("NewBook")
	err := db.BookRepository().Create(testBook)
	assert.NoError(t, err)

	book, err := db.BookRepository().FindBookByTitle("NewBook")
	assert.NoError(t, err)
	assert.NotNil(t, book)

	err = db.BookRepository().DeleteBookByTitle(book.ID)
	assert.NoError(t, err)
}
