package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/MrTj458/markednotes/postgres"
	"github.com/MrTj458/markednotes/validator"
	"github.com/MrTj458/markednotes/web"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	defer db.Close(context.Background())

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("no port provided, or it is an invalid integer")
	}

	s := web.NewServer(port)
	s.Validator = validator.New()

	s.UserService = postgres.NewUserService(db)
	s.NoteService = postgres.NewNoteService(db)
	s.FolderService = postgres.NewFolderService(db)

	log.Fatal(s.Run())
}