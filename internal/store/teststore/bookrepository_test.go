package teststore

import (
	"github.com/stretchr/testify/assert"
	"main/internal/model"
	"testing"
)

func TestBookRepository_Create(t *testing.T) {
	testStore := NewTestStore()
	testBook := NewTestBook(t)

	err := testStore.BookRepository().Create(testBook)
	assert.NoError(t, err)
}

func TestBookRepository_DeleteBookByTitle(t *testing.T) {
	t.Skip()
}

func TestBookRepository_FindBookByTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		r       bookRepository
		args    args
		want    *model.Book
		wantErr bool
	}{
		{
			name: "Right case",
			r: bookRepository{
				"Harry Potter": model.NewBook("Harry Potter"),
			},
			args:    args{"Harry Potter"},
			want:    model.NewBook("Harry Potter"),
			wantErr: false,
		},
		{
			name:    "False case",
			r:       make(bookRepository),
			args:    args{"Harry Potter"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FindBookByTitle(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBookByTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Title != tt.want.Title {
				t.Errorf("FindBookByTitle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
