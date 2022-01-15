package main

import (
	"github.com/Ad3bay0c/ddd-structure/application/handler"
	"github.com/Ad3bay0c/ddd-structure/application/server"
	"github.com/Ad3bay0c/ddd-structure/infrastructure/repository"
)

func main() {
	db := &repository.DB{}
	db.NewPostgresDB()
	app := &handler.Application{
		UserDb: db,
	}
	s := server.NewServer(app)

	s.Start()
}
