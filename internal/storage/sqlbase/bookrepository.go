package sqlbase

import (
	"github.com/jinzhu/gorm"
	"main/internal/model"
)

type bookRepository struct {
	db *gorm.DB
}

func (r *bookRepository) Create(book *model.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) FindBookByTitle(title string) (*model.Book, error) {
	book := new(model.Book)
	if err := r.db.Where("title = ?", title).First(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) DeleteBookByID(id uint) error {
	book := new(model.Book)
	if err := r.db.Where("id = ?", id).Delete(book).Error; err != nil {
		return err
	}
	return nil
}
