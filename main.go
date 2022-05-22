package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mayudev/notesplace/server/database"
	"github.com/mayudev/notesplace/server/server"
	"golang.org/x/crypto/bcrypt"
)

func parseEnv() (string, string) {
	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("warn: db user not provided. assuming postgres")
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("warn: db password not provided")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("warn: db host not provided. assuming localhost")
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("info: assuming db port 5432 (postgresql default)")
		dbPort = "5432"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("info: assuming db name notesplace")
		dbName = "notesplace"
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		fmt.Println("warn: PRIVATE KEY NOT PROVIDED. THIS IS VERY UNSAFE.")
	}

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName

	return url, privateKey
}

func main() {
	url, privateKey := parseEnv()
	store := database.NewDatabase(url)

	server := server.NewServer(store, server.ServerOptions{
		PrivateKey:  privateKey,
		HashingCost: bcrypt.DefaultCost,
	})

	if err := server.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
