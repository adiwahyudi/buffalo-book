package actions

import (
	"buffalo_book/models"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

/*
BooksShow shows the page detail books by id
Pointing to endpoint GET '/books/{id}'
*/
func BooksShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	book := models.Book{}
	bookId := c.Param("id")
	bookUuid, err := uuid.FromString(bookId)
	if err != nil {
		c.Flash().Add("warning", "wrong book id format")
		c.Render(http.StatusBadRequest, r.HTML("books/index.html"))
	}

	err = tx.Find(&book, bookUuid)
	if err != nil {
		c.Flash().Add("warning", "Book not found!")
		c.Redirect(http.StatusTemporaryRedirect, "/books")
	}

	c.Set("book", book)
	return c.Render(http.StatusOK, r.HTML("books/show.html"))
}

/*
BooksIndex shows the page list of all book.
Pointing to endpoint GET '/books'
*/
func BooksIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var books models.Books

	err := tx.All(&books)
	if err != nil {
		c.Flash().Add("warning", "Error find books")
		c.Redirect(http.StatusTemporaryRedirect, "/books")
	}

	c.Set("books", books)
	return c.Render(http.StatusOK, r.HTML("books/index.html"))
}

/*
BooksNew shows the form to create a new book.
Pointing to endpoint GET '/books/new'
*/
func BooksNew(c buffalo.Context) error {
	book := models.Book{}
	c.Set("book", book)
	return c.Render(http.StatusOK, r.HTML("books/new.html"))
}

/*
BooksCreate function to create/save a new book.
Pointing to endpoint POST '/books'
*/
func BooksCreate(c buffalo.Context) error {
	b := &models.Book{}
	if err := c.Bind(b); err != nil {
		c.Flash().Add("warning", "Form binding error")
		c.Redirect(http.StatusMovedPermanently, "/books")
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(b)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/books")
	}

	if verrs.HasAny() {
		c.Flash().Add("warning", "Form validation errors")
		c.Set("book", b)
		c.Set("errors", verrs)
		return c.Render(http.StatusBadRequest, r.HTML("books/new.html"))
	}

	c.Flash().Add("success", "Success create blog")
	return c.Redirect(http.StatusMovedPermanently, "/books")
}

/*
BooksEdit function to show edit page for books
Pointing to endpoint GET '/books/{id}/edit'
*/

func BooksEdit(c buffalo.Context) error {
	book := models.Book{}
	bookId := c.Param("id")
	bookUuid, err := uuid.FromString(bookId)
	if err != nil {
		c.Flash().Add("warning", "wrong book id format")
		c.Render(http.StatusBadRequest, r.HTML("books/index.html"))
	}

	tx := c.Value("tx").(*pop.Connection)
	err = tx.Find(&book, bookUuid)
	if err != nil {
		c.Flash().Add("info", "Book not found")
		c.Redirect(http.StatusPermanentRedirect, "/books")
	}

	c.Set("book", book)
	return c.Render(http.StatusOK, r.HTML("books/edit.html"))
}

/*
BooksUpdate function to execute update the data
Pointing to endpoint PUT '/books/{id}'
*/
func BooksUpdate(c buffalo.Context) error {
	bookId := c.Param("id")
	bookUuid, err := uuid.FromString(bookId)
	if err != nil {
		c.Flash().Add("warning", "wrong book id format")
		c.Render(http.StatusBadRequest, r.HTML("books/index.html"))
	}
	book := &models.Book{}
	if err := c.Bind(book); err != nil {
		c.Flash().Add("warning", "Form binding error")
		c.Redirect(http.StatusMovedPermanently, "/books")
	}

	book.ID = bookUuid
	tx := c.Value("tx").(*pop.Connection)
	err = tx.Update(book)
	if err != nil {
		c.Flash().Add("warning", err.Error())
		c.Redirect(http.StatusMovedPermanently, "/books")
	}

	//if verrs.HasAny() {
	//	c.Flash().Add("warning", "Form validation errors")
	//	c.Set("book", book)
	//	c.Set("errors", verrs)
	//	return c.Render(http.StatusBadRequest, r.HTML("books/edit.html"))
	//}

	c.Flash().Add("success", "Success update book")
	return c.Redirect(http.StatusMovedPermanently, "/books")
}

/*
BooksDestroy function to execute delete the data
Pointing to endpoint DELETE '/books/{id}/delete'
*/
func BooksDestroy(c buffalo.Context) error {
	c.Logger().Info("executing destroy")
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		c.Flash().Add("warning", "No transaction found")
		c.Redirect(http.StatusMovedPermanently, "/books")
	}

	bookId := c.Param("id")
	book := &models.Book{}
	err := tx.Find(book, bookId)
	if err != nil {
		c.Flash().Add("info", err.Error())
		c.Redirect(http.StatusMovedPermanently, "/books")
	}

	err = tx.Destroy(book)
	if err != nil {
		msg := fmt.Sprint("Error deleting book", book.Title, ", err: ", err.Error())
		c.Flash().Add("warning", msg)
		c.Render(http.StatusInternalServerError, r.HTML("books/index.html"))
	}

	c.Flash().Add("success", "Success deleting book")
	return c.Redirect(http.StatusMovedPermanently, "/books")
}
