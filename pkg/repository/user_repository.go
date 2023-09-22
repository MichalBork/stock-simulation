package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"stock-simulation/pkg/model"
)

type UserRepository struct {
	*SQLRepository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		SQLRepository: NewSQLRepository(db),
	}
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	rows, err := r.FindBy("users", "Username", username)
	fmt.Print(rows)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	var user model.User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
		return nil, err
	}

	fmt.Println(user)
	return &user, nil
}

func (r *UserRepository) RegisterUser(user *model.User) error {
	data := map[string]interface{}{
		"Username": user.Username,
		"Password": user.Password, // Upewnij się, że hasło jest odpowiednio zaszyfrowane przed zapisaniem!
		"Email":    user.Email,
		// Dodaj inne pola w razie potrzeby
	}

	return r.Save("users", data)
}
