package main

import (
	"log"

	"github.com/mayudev/notesplace/server/database"
	"github.com/mayudev/notesplace/server/server"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	store := database.NewDatabase("postgres://postgres:secret@localhost:5432/notesplace")

	server := server.NewServer(store, server.ServerOptions{
		PrivateKey:  "abcdef",
		HashingCost: bcrypt.DefaultCost,
	})

	if err := server.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
