package main

import (
	"log"
	"os"
	"strconv"

	"github.com/MrTj458/markednotes/postgres"
	"github.com/MrTj458/markednotes/token"
	"github.com/MrTj458/markednotes/validator"
	"github.com/MrTj458/markednotes/web"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	defer db.Close()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("no port provided, or it is an invalid integer")
	}

	signingKey := os.Getenv("SIGNING_KEY")
	if len(signingKey) == 0 {
		log.Fatal("no signing key found")
	}

	s := web.NewServer(port)

	s.Validator = validator.New()
	s.Jwt = token.New([]byte(signingKey))

	s.UserService = postgres.NewUserService(db)
	s.NoteService = postgres.NewNoteService(db)
	s.FolderService = postgres.NewFolderService(db)

	log.Fatal(s.Run())
}
