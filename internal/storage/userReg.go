package storage

import (
	"database/sql"

	"github.com/ruziba3vich/lesson222/internal/models"
)

func CreateUser(cudto CreateUserDTO, db *sql.DB) (*models.User, error) {
	query := "INSERT INTO Users(name, email) VALUES ($1, $2) RETURNING id, name, email;"
	row := db.QueryRow(query, cudto.Name, cudto.Email)
	var user models.User
	if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

type CreateUserDTO struct {
	Name  string
	Email string
}
