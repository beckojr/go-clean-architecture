package adapter

import "github.com/beckojr/go-clean-architecture/adapter/bookctlr"

// AppController app level controller
type AppController interface {
	bookctlr.BookController
}
