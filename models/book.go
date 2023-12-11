package models

import (
	"buffalo_book/utils"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// Book is used by pop to map your books database table to your go code.
type Book struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Author      string       `json:"author" db:"author"`
	Publisher   string       `json:"publisher" db:"publisher"`
	Description string       `json:"description" db:"description"`
	Year        int          `json:"year" db:"year"`
	ImageURL    string       `json:"image_url" db:"image_url"`
	ISBN        string       `json:"isbn" db:"isbn"`
	Pages       int          `json:"pages" db:"pages"`
	Stock       int          `json:"stock" db:"stock"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	ImageFile   binding.File `db:"-" form:"imageFile"`
}

// String is not required by pop and may be deleted
func (b Book) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Books is not required by pop and may be deleted
type Books []Book

// String is not required by pop and may be deleted
func (b Books) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Book) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: b.Title, Name: "Title"},
		&validators.StringIsPresent{Field: b.Author, Name: "Author"},
		&validators.StringIsPresent{Field: b.Publisher, Name: "Publisher"},
		&validators.StringIsPresent{Field: b.Description, Name: "Description"},
		&validators.IntIsPresent{Field: b.Year, Name: "Year"},
		&validators.StringIsPresent{Field: b.ImageURL, Name: "ImageURL"},
		&validators.StringIsPresent{Field: b.ISBN, Name: "Isbn"},
		&validators.IntIsPresent{Field: b.Pages, Name: "Pages"},
		&validators.IntIsPresent{Field: b.Stock, Name: "Stock"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Book) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Book) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (b *Book) BeforeValidate(tx *pop.Connection) error {
	if !b.ImageFile.Valid() {
		return errors.New("Image file not valid")
	}
	b.ImageURL = b.ImageFile.Filename
	return nil
}

func (b *Book) BeforeCreate(tx *pop.Connection) error {
	if b.ImageURL != "book_cover_default.jpg" {
		newFilename := utils.GenerateImageFilename(b.ID.String()+b.Title+b.Author+b.ImageURL) + ".jpg"
		if !strings.HasPrefix(b.ImageURL, "http") || !strings.HasPrefix(b.ImageURL, "https") {
			b.ImageURL = newFilename
		}
	}

	return nil
}
func (b *Book) AfterCreate(tx *pop.Connection) error {
	if !b.ImageFile.Valid() {
		return nil
	}
	dir := filepath.Join("./assets/images", "book_cover")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, b.ImageURL))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, b.ImageFile)
	return err
}

func (b *Book) BeforeUpdate(tx *pop.Connection) error {
	// If files upload is empty
	if b.ImageFile.Valid() {
		ext := filepath.Ext(b.ImageFile.Filename)
		newFilename := utils.GenerateImageFilename(b.ID.String()+b.Title+b.Author+b.ImageFile.Filename) + ext
		b.ImageURL = newFilename
	}

	return nil
}

func (b *Book) AfterUpdate(tx *pop.Connection) error {
	if b.ImageFile.Valid() {
		dir := filepath.Join("./assets/images", "book_cover")
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.WithStack(err)
		}
		f, err := os.Create(filepath.Join(dir, b.ImageURL))
		if err != nil {
			return errors.WithStack(err)
		}
		defer f.Close()
		_, err = io.Copy(f, b.ImageFile)
		return err

	}
	return nil
}
