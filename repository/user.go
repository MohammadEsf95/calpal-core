package repository

import (
	"calpal-core/entity"
	"database/sql"
)

type UserRepository interface {
	Users() ([]entity.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) Users() ([]entity.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []entity.User

	for rows.Next() {

		var user entity.User

		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Email,
			&user.TargetCalories,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
