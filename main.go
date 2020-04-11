package main

import (
	"log"

	"github.com/beckojr/go-clean-architecture/adapter/bookctlr"
	"github.com/beckojr/go-clean-architecture/infrastructure/database"
	"github.com/beckojr/go-clean-architecture/infrastructure/router"
	"github.com/beckojr/go-clean-architecture/usecase/booksvc"
	"github.com/labstack/echo"
)

func main() {
	db := database.NewDatabase()
	defer db.Close()
	repo := bookctlr.NewBookRepository(db)
	bookService := booksvc.New(repo)
	app := bookctlr.New(bookService)
	e := echo.New()
	e = router.New(e, app)
	log.Fatal(e.Start(":3000"))

}
