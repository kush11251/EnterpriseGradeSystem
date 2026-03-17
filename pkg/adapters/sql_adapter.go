package adapters

import (
	"database/sql"
	"fmt"

	"github.com/EnterpriseGradeSystem/pkg/models"
)

type SqlAdapter struct {}

func (s *SqlAdapter) GetUsers() ([]models.User, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}