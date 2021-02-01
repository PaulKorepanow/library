package sqlbase

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"main/internal/model"
	"strings"
	"testing"
)

func TestDatabase(t *testing.T) (*SqlStore, func(tables ...string)) {
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
