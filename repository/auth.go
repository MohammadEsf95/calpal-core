package repository

import (
	"calpal-core/entity"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthRepository interface {
	SignUp(user entity.User) (string, error)
	SignIn(req entity.User) (entity.User, error)
}

type authRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepositoryImpl {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) SignUp(signUpUser entity.User) (string, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(signUpUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	_, err = r.db.Exec("INSERT INTO users (id, first_name, last_name, username, password, email, created_at, updated_at)"+
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", signUpUser.ID, signUpUser.FirstName, signUpUser.LastName,
		signUpUser.Username, hashPassword, signUpUser.Email, time.Now(), time.Now())
	if err != nil {
		return "", err
	}

	return signUpUser.ID, nil
}

func (r *authRepositoryImpl) SignIn(req entity.User) (entity.User, error) {
	var user entity.User

	rows, err := r.db.Query("SELECT * FROM users WHERE username = $1 AND password = $2", req.Username, req.Password)
	if err != nil {
		return entity.User{}, err
	}

	rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password)

	return user, nil
}
