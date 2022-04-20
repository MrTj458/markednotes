package main

import (
	"context"
	"log"

	"github.com/MrTj458/markednotes/postgres"
	"github.com/MrTj458/markednotes/web"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close(context.Background())

	s := web.NewServer(3000)

	s.UserService = postgres.NewUserService(db)
	s.NoteService = postgres.NewNoteService(db)
	s.FolderService = postgres.NewFolderService(db)

	log.Fatal(s.Run())
}
