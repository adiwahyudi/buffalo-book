package grifts

import (
	"buffalo_book/models"

	"github.com/gobuffalo/grift/grift"
)

var _ = grift.Namespace("db", func() {

	// Add Books
	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		books := []models.Book{
			{
				Title:       "Book 1",
				Author:      "Author 1",
				Publisher:   "Publisher 1",
				Year:        2001,
				Description: "This is Book 1",
				ImageURL:    "book_cover_default.jpg",
				ISBN:        "1234567890123",
				Pages:       100,
				Stock:       100,
			},
			{
				Title:       "Book 2",
				Author:      "Author 2",
				Publisher:   "Publisher 2",
				Year:        2002,
				Description: "This is Book 2",
				ImageURL:    "book_cover_default.jpg",
				ISBN:        "1234567890223",
				Pages:       102,
				Stock:       102,
			},
			{
				Title:       "Book 3",
				Author:      "Author 3",
				Publisher:   "Publisher 3",
				Year:        2003,
				Description: "This is Book 3",
				ImageURL:    "book_cover_default.jpg",
				ISBN:        "1234567890323",
				Pages:       103,
				Stock:       103,
			},
			{
				Title:       "Book 4",
				Author:      "Author 4",
				Publisher:   "Publisher 4",
				Year:        2004,
				Description: "This is Book 4",
				ImageURL:    "book_cover_default.jpg",
				ISBN:        "1234567890423",
				Pages:       104,
				Stock:       104,
			},
			{
				Title:       "Book 5",
				Author:      "Author 5",
				Publisher:   "Publisher 5",
				Year:        2005,
				Description: "This is Book 5",
				ImageURL:    "book_cover_default.jpg",
				ISBN:        "1234567890523",
				Pages:       105,
				Stock:       105,
			},
		}

		for _, book := range books {
			err := models.DB.Create(&book)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})

})
