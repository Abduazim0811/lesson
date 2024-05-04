package storage

import (
	"database/sql"

	"github.com/ruziba3vich/lesson222/internal/models"
)

func GetAllUsers(db *sql.DB) (users []models.User, e error) {
	query := "SELECT u.id, u.name, u.email FROM Users u;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
