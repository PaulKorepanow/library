package sqlbase

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"main/internal/model"
	"strings"
	"testing"
)

func DataBase(t *testing.T) (*SqlStore, func(tables ...string)) {
	t.Helper()
	db, err := gorm.Open(
		"postgres",
		"host=localhost dbname=postgres user=postgres password=12345678 sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.AutoMigrate(&model.Book{}).Error; err != nil {
		t.Fatal(err)
	}
	return NewSqlStore(db), func(tables ...string) {
		if err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))).Error; err != nil {
			t.Fatal(err)
		}
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBookRepository_Create(t *testing.T) {
	db, close := DataBase(t)
	defer close("books")

	testBook := model.NewBook("NewBook")
	err := db.Book().Create(testBook)
	assert.NoError(t, err)
}

func TestBookRepository_FindBookByTitle(t *testing.T) {
	db, close := DataBase(t)
	defer close("books")

	testBook := model.NewBook("NewBook")
	err := db.Book().Create(testBook)
	assert.NoError(t, err)

	book, err := db.Book().FindBookByTitle("NewBook")
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, book.Title, "NewBook")
}

func TestBookRepository_DeleteBookByID(t *testing.T) {
	db, close := DataBase(t)
	defer close("books")

	testBook := model.NewBook("NewBook")
	err := db.Book().Create(testBook)
	assert.NoError(t, err)

	book, err := db.Book().FindBookByTitle("NewBook")
	assert.NoError(t, err)
	assert.NotNil(t, book)

	err = db.Book().DeleteBookByID(book.ID)
	assert.NoError(t, err)
}
