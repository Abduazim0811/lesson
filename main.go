package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ruziba3vich/lesson222/internal/storage"
)

type Customer struct {
	Id       int
	Username string
	Password string
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "lesson222")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	users, err := storage.GetAllUsers(db)

	post,err := storage.GetPosts(db)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
