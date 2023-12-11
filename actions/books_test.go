package actions

import (
	"buffalo_book/models"
	"fmt"
)

func (as *ActionSuite) Test_Books_Show() {
	as.LoadFixture("books scenario")

	book := models.Book{}

	err := as.DB.Last(&book)
	if err != nil {
		panic(err)
	}

	res := as.HTML(fmt.Sprintf("/books/%s", book.ID)).Get()
	body := res.Body.String()
	as.Contains(body, "Book 5", "Book name appears on Show page.")
}

func (as *ActionSuite) Test_Books_Index() {
	as.LoadFixture("books scenario")
	book := models.Books{}

	totalBook, err := as.DB.Count(&book)
	if err != nil {
		panic(err)
	}

	res := as.HTML("/books").Get()
	body := res.Body.String()

	as.Contains(body, "All Book", "Page title appears on Index page")
	as.Contains(body, "Book 1", "Book 1 appears on Index page")
	as.Contains(body, "Book 5", "Book 5 appears on Index page")
	as.Equal(5, totalBook, "Total book loaded from fixture should be 5 ")
}
