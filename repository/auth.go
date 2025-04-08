package repository

import (
	"calpal-core/entity"
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthRepository interface {
	SignUp(user entity.User) (string, error)
	SignIn() (entity.User, error)
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

	signUpUser.Password = string(hashPassword)
	signUpUser.ID = uuid.New().String()

	_, err = r.db.Exec("INSERT INTO users (id, first_name, last_name, password, email, created_at, updated_at)"+
		" VALUES ($1, $2, $3, $4, $5, $6, $7)", signUpUser.ID, signUpUser.FirstName, signUpUser.LastName, signUpUser.Password,
		signUpUser.Email, time.Now(), time.Now())
	if err != nil {
		return "", err
	}

	return signUpUser.ID, nil
}

func (r *authRepositoryImpl) SignIn() (entity.User, error) {
	panic("implement me")
}
